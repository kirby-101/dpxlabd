package routes

import (
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"net/http"
	"slices"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

func ObfuscateRoutes[T any](routes map[string]T) map[string]T {
	var obfusRoutes = make(map[string]T, len(routes))

	for route, h := range routes {
		var (
			sum         = sha256.Sum256([]byte(route))
			obfus_route = "/" + base64.RawURLEncoding.EncodeToString(
				sum[:8], //?
			)
		)

		obfusRoutes[obfus_route] = h
	}

	return obfusRoutes
}

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
