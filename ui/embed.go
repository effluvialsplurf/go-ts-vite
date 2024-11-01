package ui

import (
	"embed"
)

//go:embed frontend/dist/*
//go:embed frontend/dist/index.html
var Index embed.FS
