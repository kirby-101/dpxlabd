package routes

import (
	"bytes"
	"embed"
	"html/template"
	"path/filepath"
)

const (
	basePage  = "base.html"
	errorPage = "error.html"
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
// returns rendered html, error
func renderHTMLTemplate(templateFilePath string, templateData any) (string, error) {
	t, err := template.New(filepath.Base(templateFilePath)).
		Funcs(template.FuncMap{
			"safeHTML": func(s template.HTML) template.HTML { return s },
		}).
		ParseFiles(templateFilePath)
	if err != nil {
		return "", err
	}

	var buf = &bytes.Buffer{}
	if err := t.ExecuteTemplate(buf, filepath.Base(templateFilePath), templateData); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func renderMarkdownToHTML(_markdown []byte) string {
	var (
		p   = parser.New()
		doc = p.Parse(_markdown)
		r   = html.NewRenderer(
			html.RendererOptions{
				Flags: html.CommonFlags,
			},
		)
	)

	return string(markdown.Render(doc, r))
}

// returns map[urlPath]fsPath
func genFsMap(directory, urlPrefix string) (map[string]string, error) {
	var pathMap = map[string]string{}

	// walk & store
	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walkFunc lastErr: %s: %s", path, err.Error())
		}

		if !info.IsDir() {
			// relative path
			relativePath, err := filepath.Rel(directory, path)
			if err != nil {
				return fmt.Errorf("walkFunc curErr: %s: %s", path, err.Error())
			}

			uriPath := urlPrefix + "/" + relativePath
			pathMap[uriPath] = path
		}

		return nil
	}

	if err := filepath.Walk(
		directory,
		walkFunc,
	); err != nil {
		return nil, err
	}

	return pathMap, nil
}
*/
