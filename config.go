package yeppeun

import "embed"

//go:embed templates/**/*.tmpl
var VFS embed.FS
