package md

import (
	"embed"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

//go:embed content/*
var markdownFS embed.FS

func catServeMarkdownFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	routes.CatServeFS(ctx, markdownFS, filePath)
}
