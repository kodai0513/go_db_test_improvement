// Package migrations embeds the goose migration set so that both the
// application (cmd/api) and every infrastructure test package can replay
// the exact same schema without depending on a filesystem path.
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
