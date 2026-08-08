// Package testhelper provides the PostgreSQL testcontainer bootstrap that
// every internal/infrastructure/*/postgres package uses in its own
// TestMain.
//
// This function is intentionally shared/DRY: the bug this repository
// reproduces is NOT code duplication. It is that TestMain runs once per Go
// *package* (= once per compiled test binary), and `go test ./...` runs
// package test binaries concurrently (up to GOMAXPROCS). So even though
// every package calls the exact same StartPostgres helper, running the
// whole suite still boots one full PostgreSQL container - and replays the
// entire goose migration set - per package, all at once.
package testhelper

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"

	"example.com/go-db-test-improvement/db/migrations"
)

// StartPostgres boots a fresh postgres:16-alpine container, replays the
// full goose migration set inside it, and returns a ready-to-use *sql.DB
// together with a cleanup function that closes the connection and
// terminates the container.
func StartPostgres() (*sql.DB, func(), error) {
	ctx := context.Background()

	container, err := tcpostgres.Run(ctx,
		"postgres:16-alpine",
		tcpostgres.WithDatabase("app"),
		tcpostgres.WithUsername("app"),
		tcpostgres.WithPassword("app"),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("start postgres container: %w", err)
	}
	cleanup := func() {
		_ = container.Terminate(context.Background())
	}

	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("connection string: %w", err)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("open db: %w", err)
	}

	// Postgres restarts itself once during container init (initdb, then a
	// real startup), and under docker-outside-of-docker the container's
	// readiness probe races that restart: the published port can accept a
	// TCP connection or report "starting up" for a brief window after
	// testcontainers already reported the container ready. Under heavy
	// concurrent load (many sibling Postgres containers + go test binaries
	// competing for CPU) that window can stretch well past a few seconds,
	// so retry the first ping generously instead of failing on it.
	var pingErr error
	for i := 0; i < 300; i++ {
		if pingErr = db.PingContext(ctx); pingErr == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	if pingErr != nil {
		cleanup()
		return nil, nil, fmt.Errorf("ping db: %w", pingErr)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("set goose dialect: %w", err)
	}
	goose.SetBaseFS(migrations.FS)
	if err := goose.Up(db, "."); err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("run migrations: %w", err)
	}

	return db, func() {
		_ = db.Close()
		cleanup()
	}, nil
}
