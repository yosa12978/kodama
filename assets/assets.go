package assets

import (
	"embed"
)

//go:embed assets/*
var AssetsFS embed.FS
