#!/usr/bin/env bash
# Polls `docker stats` every 2s and appends CSV rows so the CPU spike
# during scripts/repro.sh can be reviewed afterwards. Run this in a
# separate terminal, on the host, before starting the reproduction.
set -euo pipefail

OUT="${1:-docker-stats.log}"
echo "timestamp,name,cpu_perc,mem_usage,mem_perc" > "$OUT"

echo "Watching docker stats -> $OUT (Ctrl+C to stop)"
while true; do
  ts=$(date -u +%Y-%m-%dT%H:%M:%SZ)
  docker stats --no-stream --format '{{.Name}},{{.CPUPerc}},{{.MemUsage}},{{.MemPerc}}' \
    | sed "s/^/${ts},/" >> "$OUT"
  sleep 2
done
