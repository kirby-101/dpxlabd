package assets

import (
	"embed"
	"path"
	"strings"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

//go:embed content/*
var assetsFS embed.FS

func catServeAssetsFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	filePath = strings.TrimPrefix(filePath, "/assets/")

	routes.CatServeFS(
		ctx,
		assetsFS,
		path.Join("content", filePath),
	)
}
