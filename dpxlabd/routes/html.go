package routes

import (
	"bytes"
	"embed"
	"html/template"
	"path/filepath"
)

const (
	basePage  = "base.html"
	errorPage = "html/error.html"
	indexPage = "index.html"
)

//go:embed html/*
var htmlFS embed.FS

func RenderHTML(target string, data any) (string, error) {
	var (
		t = template.Must(
			template.New(
				filepath.Base(target),
			).ParseFS(htmlFS, target),
		)
		buf = &bytes.Buffer{}
	)

	return buf.String(), t.Execute(buf, data)
}

func RenderErrorHTML(errCode, errMesg string) (string, error) {
	return RenderHTML(
		errorPage,
		struct {
			ErrorCode    string
			ErrorMessage string
		}{
			ErrorCode:    errCode,
			ErrorMessage: errMesg,
		},
	)
}

/*
80
[2026-09-27T04:30:56.810Z]-[error]-[github.com/kirby-101/dpxlab/dpxlabd/routes.(*ErrorRouteHandler).Handle] handler probably not found
2026/09/27 04:30:56 http: panic serving 10.0.0.1:52837: template: pattern matches no files: `error.html`
goroutine 19 [running]:
net/http.(*conn).serve.func1()
        net/http/server.go:1897 +0xbd
panic({0x821280?, 0x2302ea094190?})
        runtime/panic.go:860 +0x13a
html/template.Must(...)
        html/template/template.go:368
github.com/kirby-101/dpxlab/dpxlabd/routes.RenderHTML({0x89c9aa, 0xa}, {0x83f280, 0x2302ea0b8100})
        github.com/kirby-101/dpxlab/dpxlabd/routes/html.go:21 +0x131
github.com/kirby-101/dpxlab/dpxlabd/routes.RenderErrorHTML(...)
        github.com/kirby-101/dpxlab/dpxlabd/routes/html.go:33
github.com/kirby-101/dpxlab/dpxlabd/routes.Error(0x2302ea0b80e0, 0x194)
        github.com/kirby-101/dpxlab/dpxlabd/routes/errors.go:75 +0x185
github.com/kirby-101/dpxlab/dpxlabd/routes.(*ErrorRouteHandler).Handle(0x0?, 0x2302ea0b80e0)
        github.com/kirby-101/dpxlab/dpxlabd/routes/errors.go:53 +0x54
github.com/kirby-101/dpxlab/dpxlabd.(*Orchestrator).HandleRoute(0x2302e9fba500, {0x8ec540, 0x2302ea09a1e0}, 0x2302ea0be000)
        github.com/kirby-101/dpxlab/dpxlabd/www.go:29 +0xfa
net/http.HandlerFunc.ServeHTTP(0x2302e9fe40c0?, {0x8ec540?, 0x2302ea09a1e0?}, 0x6f9096?)
        net/http/server.go:2286 +0x29
net/http.(*ServeMux).ServeHTTP(0x47ff99?, {0x8ec540, 0x2302ea09a1e0}, 0x2302ea0be000)
        net/http/server.go:2828 +0x1c7
net/http.serverHandler.ServeHTTP({0x2302ea0b6080?}, {0x8ec540?, 0x2302ea09a1e0?}, 0x6?)
        net/http/server.go:3311 +0x8e
net/http.(*conn).serve(0x2302ea0ba090, {0x8ecfa8, 0x2302ea09c240})
        net/http/server.go:2067 +0x690
created by net/http.(*Server).Serve in goroutine 18
        net/http/server.go:3464 +0x485
*/
