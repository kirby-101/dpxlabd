package routes

import "net/http"

func SetContentType(value string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", value)
}

func SetContentTypeHTML(w http.ResponseWriter) {
	SetContentType("text/html; charset=utf-8", w)
}

func SetContentTypeText(w http.ResponseWriter) {
	SetContentType("text/plain; charset=utf-8", w)
}

func DoNotCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
}

/*
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
*/
