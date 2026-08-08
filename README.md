# go-db-test-improvement

`testcontainers-go` を使ったDBテストが、パッケージ数の増加とともにCPUを食い潰していく現象を、実際に手を動かして再現するためのサンプルプロジェクトです。

## 再現したい現象

Go + Echo + oapi-codegen + DDD + sqlc + goose という構成のアプリケーションでは、インフラ層を機能(境界づけられたコンテキスト)ごとにパッケージ分割することが多くあります。

```
internal/infrastructure/
├── member/postgres/       (会員)
├── product/postgres/      (商品)
├── inventory/postgres/    (在庫)
├── order/postgres/        (受注)
├── payment/postgres/      (決済)
├── shipment/postgres/     (配送)
├── coupon/postgres/       (クーポン)
├── review/postgres/       (レビュー)
└── notification/postgres/ (通知)
```

このリポジトリでは、上記9つのパッケージそれぞれが**独自の `TestMain`** を持ち、共通の `internal/testhelper.StartPostgres()` を呼び出して自分専用の PostgreSQL コンテナを起動 → goose migrationをフルセット再生 → テスト実行、という作りになっています(`internal/infrastructure/*/postgres/repository_test.go`)。

`StartPostgres` 自体はコードとして共通化されているにもかかわらず、`TestMain` は **Goのパッケージ単位(=テストバイナリ単位)** に存在するため、`go test ./...` を実行すると:

- パッケージごとに新しいテストバイナリが起動する
- `go test` はデフォルトで `-p`(≒GOMAXPROCS)個のパッケージを並列実行する
- 結果として、9個のPostgresコンテナがほぼ同時に起動し、9セット分のgoose migrationが同時に走る

というシナリオが発生し、CPU使用率が跳ね上がります。「コードを共通化していても、コンテナのライフサイクルがパッケージ単位にスコープされている限りボトルネックは解消しない」という点が、この現象の本質です。

## 構成

- **ドメイン**: ECサイト基幹システム(会員/商品/在庫/受注/決済/配送/クーポン/レビュー/通知の9境界コンテキスト)
- **db/migrations**: goose migration。全コンテキスト分のテーブルを1セットで管理し、`//go:embed` で埋め込み
- **db/queries, sqlc.yaml**: sqlcの入力。生成結果は `internal/infrastructure/*/postgres/sqlc/` に手動で反映済み(このリポジトリには sqlc CLI 実行環境がないため、生成後のコードを直接コミットしています。`sqlc generate` を実行すればこの内容が再生成されます)
- **internal/domain/**: ドメイン層(エンティティ + リポジトリIF)
- **internal/infrastructure/**: ★ボトルネック再現地点。各パッケージが独自 `TestMain` を持つ
- **internal/testhelper**: 全パッケージ共通のPostgresコンテナ起動処理
- **cmd/api, internal/api, openapi/**: Echo + oapi-codegen 構成の最小サンプル(会員APIのみ実装。他8コンテキストも同一パターンで追加できます)
- **docker-compose.yml / Dockerfile.testrunner**: 誰の環境でも同じ条件で再現できるように、`go test ./...` をコンテナ内で実行し、ホストのDockerデーモンに対して testcontainers 経由でPostgresコンテナを起動させる構成(docker-outside-of-docker)
- **docker-compose.dev.yml**: `cmd/api` をローカルで動かす際の開発用DB(ボトルネック再現とは無関係)

## 再現手順

必要なもの: Docker Desktop(または dockerd + docker compose v2)。ホストにGoをインストールする必要はありません。

```bash
# 1. 別ターミナルでCPU/メモリの推移を記録し始める
bash scripts/watch-stats.sh
# Windows PowerShellの場合: ./scripts/watch-stats.ps1

# 2. 本体: go test ./... をコンテナ内で実行
docker compose build test-runner
docker compose run --rm test-runner
# または: make repro
```

`docker-stats.log` に、9個の `postgres:16-alpine` コンテナがほぼ同時に立ち上がり、CPU使用率が積み上がっていく様子がCSVで記録されます。ホストのコア数が少ないマシンほど、オーバーサブスクリプションによる悪化が顕著に見えるはずです(`Dockerfile.testrunner` は `-p 9` を明示指定しているため、コア数に関わらず9パッケージが強制的に同時実行されます)。

## 注意点

- このリポジトリは **問題を再現するためのもの** で、改善版(コンテナ共有・並列度制御など)はまだ含んでいません。
- `go.sum` はコミットしていません。`Dockerfile.testrunner` がビルド時に `go mod tidy` を実行して解決します(ビルド時にネットワークアクセスが必要です)。ローカルでGoを使う場合は先に `make tidy` を実行してください。
- `internal/infrastructure/*/postgres/sqlc/*.go` は sqlc の生成コードを模して手書きしたものです。`sqlc generate` を実際に実行すればこれと同等のコードが再生成される想定です。
- この環境(本セッションのサンドボックス)には Go や Docker が入っていなかったため、`go build` / `go test` / `docker compose` の実行確認はできていません。お手元の環境で `make repro` を実行して動作を確認してください。
