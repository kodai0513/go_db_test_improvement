// Package testhelper は、15個の internal/infrastructure/*/postgres パッケージが
// それぞれの TestMain から呼び出す、共有Postgresコンテナ + テンプレートDBの
// ブートストラップを提供する。
//
// アーキテクチャ: testcontainers-go の Reuse 機構により、最大15個の並行する
// go test プロセス全員が同じ名前のコンテナを取り合い、最初の1プロセスだけが
// 実際にコンテナを作成し、残りは自動的にそれへアタッチする(ライブラリが
// プロセス間の名前衝突を内部でハンドリングする)。コンテナは tmpfs 上にデータを
// 持ち、Ryuk reaper が「このセッションの全接続が切れたら」自動的に破棄する。
//
// マイグレーションは "test_template" という1つのテンプレートDBに対して
// セッションレベルのアドバイザリロックで直列化しながら1回だけ実行し、
// 各パッケージはそのテンプレートを CREATE DATABASE ... WITH TEMPLATE で
// 高速クローンした自分専用のDBを使う(150マイグレーションの再生をパッケージ数分
// 繰り返すコストをなくす)。
package testhelper

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"

	"example.com/go-db-test-improvement/db/migrations"
)

const (
	// sharedContainerName は全パッケージが取り合う共有コンテナの固定名。
	// testcontainers.WithReuseByName に渡すことで、複数プロセスが同時に
	// Run() を呼んでも1個しか作られないことをライブラリが保証する。
	sharedContainerName = "go-db-test-improvement-shared-postgres"

	// templateDBName はプロセス間で1回だけマイグレーションを流す対象。
	templateDBName = "test_template"

	// advisoryLockKey はテンプレート構築を直列化するための固定キー。
	advisoryLockKey = 78347201

	pingAttempts = 30
	pingInterval = 200 * time.Millisecond
)

// StartPostgres は共有Postgresコンテナ(無ければ作成、あれば再利用)へ接続し、
// マイグレーション済みテンプレートDBを(未構築なら)構築したうえで、
// 呼び出し元パッケージ専用のDBをテンプレートから複製して返す。
func StartPostgres() (db *sql.DB, cleanup func(), err error) {
	ctx := context.Background()

	container, err := tcpostgres.Run(ctx,
		"postgres:16-alpine",
		tcpostgres.WithDatabase("app"),
		tcpostgres.WithUsername("app"),
		tcpostgres.WithPassword("app"),
		testcontainers.WithReuseByName(sharedContainerName),
		testcontainers.WithTmpfs(map[string]string{"/var/lib/postgresql/data": "rw"}),
		testcontainers.WithEnv(map[string]string{
			// tmpfsを /var/lib/postgresql/data に直接マウントすると、
			// postgres:16-alpineのinitdbが「ディレクトリが空でない」と
			// 誤検知することがあるため、サブディレクトリにPGDATAをずらす。
			"PGDATA": "/var/lib/postgresql/data/pgdata",
		}),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("start/attach shared postgres container: %w", err)
	}
	// 注意: ここで container.Terminate() は呼ばない。
	// 他の並行実行パッケージがまだこのコンテナを使っている可能性があるため、
	// コンテナ自体の破棄はRyuk reaperに委ねる(全プロセス終了時に自動実行)。

	adminDSN, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, nil, fmt.Errorf("connection string: %w", err)
	}

	if err := ensureTemplate(ctx, adminDSN); err != nil {
		return nil, nil, err
	}

	// os.Getpid() は同一 go test ./... 実行内の並行プロセス間ではユニークだが、
	// コンテナをReuseしたまま複数回 docker compose run を繰り返すと
	// OSがPIDを再利用しうるため、ナノ秒タイムスタンプを足して衝突を防ぐ。
	cloneName := fmt.Sprintf("test_%d_%d", os.Getpid(), time.Now().UnixNano())
	if err := cloneTemplate(ctx, adminDSN, cloneName); err != nil {
		return nil, nil, err
	}

	cloneDSN, err := withDatabase(adminDSN, cloneName)
	if err != nil {
		dropDatabase(context.Background(), adminDSN, cloneName)
		return nil, nil, fmt.Errorf("build clone dsn: %w", err)
	}

	db, err = sql.Open("pgx", cloneDSN)
	if err != nil {
		dropDatabase(context.Background(), adminDSN, cloneName)
		return nil, nil, fmt.Errorf("open clone db: %w", err)
	}
	if err := waitForPing(ctx, db); err != nil {
		db.Close()
		dropDatabase(context.Background(), adminDSN, cloneName)
		return nil, nil, fmt.Errorf("ping clone db: %w", err)
	}

	cleanup = func() {
		_ = db.Close()
		dropDatabase(context.Background(), adminDSN, cloneName)
	}
	return db, cleanup, nil
}

// ensureTemplate は templateDBName が存在し、マイグレーション済みで、
// 以後の接続を拒否する状態(sealed)になっていることを保証する。
// 複数プロセス・ゴルーチンから同時に呼ばれても構築が1回だけになるよう、
// Postgresのセッションレベルアドバイザリロックで直列化する。
func ensureTemplate(ctx context.Context, adminDSN string) error {
	adminPool, err := sql.Open("pgx", adminDSN)
	if err != nil {
		return fmt.Errorf("open admin db: %w", err)
	}
	defer adminPool.Close()
	if err := waitForPing(ctx, adminPool); err != nil {
		return fmt.Errorf("ping admin db: %w", err)
	}

	// pg_advisory_lock はセッション単位のロックなので、ロック取得から解放まで
	// 必ず同一の物理コネクション(*sql.Conn)を使い続ける必要がある。
	// プールされた*sql.DBに直接クエリを投げると、呼び出しごとに別バックエンドへ
	// 割り振られてしまい、ロックの意味がなくなる。
	conn, err := adminPool.Conn(ctx)
	if err != nil {
		return fmt.Errorf("acquire admin conn: %w", err)
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, "SELECT pg_advisory_lock($1)", advisoryLockKey); err != nil {
		return fmt.Errorf("acquire advisory lock: %w", err)
	}
	defer func() {
		_, _ = conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock($1)", advisoryLockKey)
	}()

	sealed, err := templateSealed(ctx, conn)
	if err != nil {
		return fmt.Errorf("check template sealed: %w", err)
	}
	if sealed {
		// 他のプロセスが既に構築済み。何もせず抜ける。
		return nil
	}

	exists, err := templateExists(ctx, conn)
	if err != nil {
		return fmt.Errorf("check template existence: %w", err)
	}
	if !exists {
		if _, err := conn.ExecContext(ctx, "CREATE DATABASE "+quoteIdent(templateDBName)); err != nil {
			return fmt.Errorf("create template db: %w", err)
		}
	}

	templateDSN, err := withDatabase(adminDSN, templateDBName)
	if err != nil {
		return fmt.Errorf("build template dsn: %w", err)
	}
	// migrateTemplate は自分専用のコネクションプールを開いて完全に閉じてから
	// returnする。CREATE DATABASE ... WITH TEMPLATE はクローン元への接続が
	// ゼロであることを要求するため、sealする前にこの呼び出しが完了し、
	// テンプレートへの接続が確実にゼロになっている必要がある。
	if err := migrateTemplate(ctx, templateDSN); err != nil {
		return fmt.Errorf("migrate template: %w", err)
	}

	if _, err := conn.ExecContext(ctx, "ALTER DATABASE "+quoteIdent(templateDBName)+" WITH is_template = true"); err != nil {
		return fmt.Errorf("mark template: %w", err)
	}
	if _, err := conn.ExecContext(ctx, "ALTER DATABASE "+quoteIdent(templateDBName)+" WITH ALLOW_CONNECTIONS false"); err != nil {
		return fmt.Errorf("seal template: %w", err)
	}
	return nil
}

func templateExists(ctx context.Context, conn *sql.Conn) (bool, error) {
	var exists bool
	err := conn.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", templateDBName,
	).Scan(&exists)
	return exists, err
}

func templateSealed(ctx context.Context, conn *sql.Conn) (bool, error) {
	var allowConn bool
	err := conn.QueryRowContext(ctx,
		"SELECT datallowconn FROM pg_database WHERE datname = $1", templateDBName,
	).Scan(&allowConn)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return !allowConn, nil
}

// migrateTemplate は goose の全マイグレーションを templateDSN に対して実行する。
// goose_db_version による追跡のおかげで、既に適用済みのマイグレーションは
// 再実行してもスキップされる(クラッシュ後の再構築でも安全)。
func migrateTemplate(ctx context.Context, templateDSN string) error {
	db, err := sql.Open("pgx", templateDSN)
	if err != nil {
		return fmt.Errorf("open template db: %w", err)
	}
	defer db.Close()

	if err := waitForPing(ctx, db); err != nil {
		return fmt.Errorf("ping template db: %w", err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}
	goose.SetBaseFS(migrations.FS)
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}

// cloneTemplate はテンプレートDBには一切接続せず、admin(app)DB経由で
// CREATE DATABASE ... WITH TEMPLATE を発行してパッケージ専用DBを複製する。
func cloneTemplate(ctx context.Context, adminDSN, name string) error {
	adminPool, err := sql.Open("pgx", adminDSN)
	if err != nil {
		return fmt.Errorf("open admin db: %w", err)
	}
	defer adminPool.Close()

	stmt := fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s", quoteIdent(name), quoteIdent(templateDBName))
	if _, err := adminPool.ExecContext(ctx, stmt); err != nil {
		return fmt.Errorf("clone template into %s: %w", name, err)
	}
	return nil
}

// dropDatabase は呼び出し元パッケージ専用のクローンDBだけを消す。
// テンプレートDBや共有コンテナ自体には一切触れない。
func dropDatabase(ctx context.Context, adminDSN, name string) {
	adminPool, err := sql.Open("pgx", adminDSN)
	if err != nil {
		return
	}
	defer adminPool.Close()
	// WITH (FORCE) (PG13+) は残存コネクションを強制切断してからDROPするため、
	// db.Close()の完了タイミングと多少ずれても確実に片付く。
	_, _ = adminPool.ExecContext(ctx, "DROP DATABASE IF EXISTS "+quoteIdent(name)+" WITH (FORCE)")
}

func withDatabase(dsn, name string) (string, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}
	u.Path = "/" + name
	return u.String(), nil
}

func quoteIdent(name string) string {
	return pgx.Identifier{name}.Sanitize()
}

func waitForPing(ctx context.Context, db *sql.DB) error {
	var err error
	for range pingAttempts {
		if err = db.PingContext(ctx); err == nil {
			return nil
		}
		time.Sleep(pingInterval)
	}
	return err
}
