package md

import (
	"embed"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

const markdownContentType = "text/html; charset=utf-8"

//go:embed content/*
var markdownFS embed.FS

func catServeMarkdownFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	routes.CatServeFS(ctx, markdownFS, filePath)
}
