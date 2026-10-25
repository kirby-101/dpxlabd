package routes

import (
	"bytes"
	"embed"
	"html/template"
	"path/filepath"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

const (
	HTMLContentType = "text/html; charset=utf-8"

	BasePagePath  = "html/base.html"
	ErrorPagePath = "html/error.html"
	IndexPagePath = "html/index.html"
	LoginPagePath = "html/login.html"
)

//go:embed html/*
var htmlFS embed.FS

func CatServeHTMLFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	ctx.W.Header().Set("Content-Type", HTMLContentType)
	CatServeFS(ctx, htmlFS, filePath)
}

func RenderHTML(targetPage string, data any) ([]byte, error) {
	var (
		t = template.Must(
			template.New(
				filepath.Base(targetPage),
			).ParseFS(htmlFS, targetPage),
		)
		buf = &bytes.Buffer{}
		err = t.Execute(buf, data)
	)

	return buf.Bytes(), err
}

func RenderErrorHTML(errCode, errMesg string) ([]byte, error) {
	return RenderHTML(
		ErrorPagePath,
		struct {
			ErrorCode    string
			ErrorMessage string
		}{
			ErrorCode:    errCode,
			ErrorMessage: errMesg,
		},
	)
}
