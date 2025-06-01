package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html.tmpl
var templates embed.FS

const (
	layoutFile = "layout.html.tmpl"
)

var IndexTemplate = template.Must(
	template.ParseFS(templates, layoutFile, "index.html.tmpl"),
)

var ErrorTemplate = template.Must(
	template.ParseFS(templates, layoutFile, "error.html.tmpl"),
)
