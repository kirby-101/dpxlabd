package routes

import (
	"bytes"
	"html/template"
	"path/filepath"
)

const (
	HTMLPagesPath     = "assets/html"
	HTMLPagePathBase  = HTMLPagesPath + "/base.html"
	HTMLPagePathError = HTMLPagesPath + "/error.html"
	HTMLPagePathIndex = HTMLPagesPath + "/index.html"
	HTMLPagePathLogin = HTMLPagesPath + "/login.html"

	MarkdownPagesPath = "assets/markdown"
	StaticFilesPath   = "assets/static"
)

func RenderHTML(targetPage string, data any) ([]byte, error) {
	var (
		t = template.Must(
			template.New(
				filepath.Base(targetPage),
			).ParseFS(AssetsFS, targetPage),
		)
		buf = &bytes.Buffer{}
		err = t.Execute(buf, data)
	)

	return buf.Bytes(), err
}

func RenderErrorHTML(errCode, errMesg string) ([]byte, error) {
	return RenderHTML(
		HTMLPagePathError,
		struct {
			ErrorCode    string
			ErrorMessage string
		}{
			ErrorCode:    errCode,
			ErrorMessage: errMesg,
		},
	)
}
