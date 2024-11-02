package ui

import (
	"embed"
)

//go:embed staticPages/*
//go:embed frontend/dist/*
var Index embed.FS
