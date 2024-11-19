package public

import "embed"

//go:embed all:template
var TemplateDir embed.FS

//go:embed all:static
var StaticDir embed.FS
