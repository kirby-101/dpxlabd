package routes

import (
	"embed"
	"net/http"
	"slices"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

// serves pages from a embed.FS
func CatServeFS(ctx *dpxlabd.RouteHandlerContext, fs embed.FS, filePath string) {
	buf, err := fs.ReadFile(filePath)
	if err != nil {
		Error(ctx, http.StatusInternalServerError)

		return
	}

	ctx.W.Write(buf)
}

// func(ctx, []http.Method*) => bool:isOkay?
// renders errPage if !isOkay
func AssertHttpMethod(ctx *dpxlabd.RouteHandlerContext, whitelist []string) bool {
	if !slices.Contains(whitelist, ctx.R.Method) {
		Error(ctx, http.StatusMethodNotAllowed)

		return false
	}

	return true
}
