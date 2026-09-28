package md

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type MarkdownRoute struct {
	AllowedMethods []string
	AllowedRoutes  map[string][]byte
}

// ! convert md to html
func NewMarkdownRoute(routes map[string]dpxlabd.RouteHandler) *MarkdownRoute {
	var r = &MarkdownRoute{
		AllowedMethods: []string{http.MethodGet},
		AllowedRoutes:  map[string][]byte{},
	}

	err := fs.WalkDir(
		markdownFS,
		"content",
		func(path string, dir fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if dir.IsDir() {
				return nil
			}

			// content/asset.file => /assets/asset.file
			var routerPath = "/assets/" + strings.TrimPrefix(path, "content/")

			routes[routerPath] = r
			r.AllowedRoutes[routerPath] = []byte{} //! content

			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	return r
}

func (r *MarkdownRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	var route = ctx.R.URL.Path

	// is valid route?
	if _, oke := r.AllowedRoutes[route]; !oke {
		routes.Error(ctx, http.StatusNotFound)

		return
	}

	// serve markdown

	//catServeMarkdownFS(ctx, ctx.R.URL.Path)
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
