package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html.tmpl
var templates embed.FS

var IndexTemplate = template.Must(
	template.ParseFS(templates, "layout.html.tmpl", "index.html.tmpl"),
)
