package routes

import (
	"bytes"
	"embed"
	"html/template"
	"path/filepath"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

const (
	htmlContentType = "text/html; charset=utf-8"

	basePage  = "html/base.html"
	errorPage = "error.html"
	indexPage = "html/index.html"
)

//go:embed html/*
var htmlFS embed.FS

func CatServeHTMLFS(ctx *dpxlabd.RouteHandlerContext, filePath string) {
	//ctx.W.Header().Set("Content-Type", htmlContentType)
}

func RenderHTML(targetPage string, data any) (string, error) {
	var (
		t = template.Must(
			template.New(
				filepath.Base(targetPage),
			).ParseFiles(targetPage),
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
