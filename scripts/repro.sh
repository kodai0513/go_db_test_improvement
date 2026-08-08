#!/usr/bin/env bash
# Reproduces the testcontainer-per-package CPU bottleneck:
# builds and runs `go test ./...` inside a container that talks to the
# host's Docker daemon, so all 9 infrastructure packages boot their own
# Postgres container at roughly the same time.
#
# Run scripts/watch-stats.sh (or .ps1) in another terminal first to see the
# CPU/memory spike as it happens.
set -euo pipefail
cd "$(dirname "$0")/.."

docker compose build test-runner
docker compose run --rm test-runner
