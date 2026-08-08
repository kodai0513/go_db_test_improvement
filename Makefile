.PHONY: repro watch dev-up dev-down tidy test-local

## Reproduce the testcontainer-per-package CPU bottleneck (docker + docker compose).
repro:
	docker compose build test-runner
	docker compose run --rm test-runner

## Run in another terminal while `make repro` is running.
watch:
	bash scripts/watch-stats.sh

## Local dev DB for cmd/api (unrelated to the bottleneck reproduction).
dev-up:
	docker compose -f docker-compose.dev.yml up -d

dev-down:
	docker compose -f docker-compose.dev.yml down -v

## Resolve go.sum (needs network + a local Go toolchain).
tidy:
	go mod tidy

## Run the same `go test ./...` directly on the host instead of in a
## container - only meaningful if you have Go + Docker installed locally.
test-local:
	go test ./... -v -p 9
