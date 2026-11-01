package routes

import "embed"

//go:embed assets
var AssetsFS embed.FS

func ServeHTMLPage() {}
