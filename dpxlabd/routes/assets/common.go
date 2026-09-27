package assets

import (
	"embed"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

//go:embed content/*
var assetsFS embed.FS

func catServeAssetsFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	routes.CatServeFS(ctx, assetsFS, filePath)
}
