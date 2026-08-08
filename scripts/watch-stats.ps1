# Polls `docker stats` every 2s and appends CSV rows so the CPU spike
# during scripts/repro.sh (or docker compose run --rm test-runner) can be
# reviewed afterwards. Run this in a separate PowerShell window, on the
# host, before starting the reproduction.
param(
    [string]$OutFile = "docker-stats.log"
)

"timestamp,name,cpu_perc,mem_usage,mem_perc" | Out-File -FilePath $OutFile -Encoding utf8

Write-Host "Watching docker stats -> $OutFile (Ctrl+C to stop)"
while ($true) {
    $ts = [DateTime]::UtcNow.ToString("yyyy-MM-ddTHH:mm:ssZ")
    $lines = docker stats --no-stream --format '{{.Name}},{{.CPUPerc}},{{.MemUsage}},{{.MemPerc}}'
    foreach ($line in $lines) {
        "$ts,$line" | Out-File -FilePath $OutFile -Append -Encoding utf8
    }
    Start-Sleep -Seconds 2
}
