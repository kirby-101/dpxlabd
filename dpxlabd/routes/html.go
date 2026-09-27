package routes

import (
	"bytes"
	"embed"
	"html/template"
	"path/filepath"
)

const (
	basePage  = "html/base.html"
	errorPage = "html/error.html"
	indexPage = "html/index.html"
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
