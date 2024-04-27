package views

import (
	"html/template"
	"io"
	"io/fs"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"
)

type (
    // FuncMapper provides a [template.FuncMap] from an [echo.Context].
	FuncMapper func(echo.Context) template.FuncMap

    // Renderer provides a custom implementation of [echo.Renderer] using
    // [template.Template].
	Renderer struct {
		tmpl   *template.Template
		mapper FuncMapper
		dirfs  fs.FS
	}

    // Options configure the Renderer.
	Options struct {
        // The root directory of the templates.
		RootDir     string
        // The glob pattern to match when initializing.
		GlobPattern string
        // Optional list of [template.FuncMap] objects.
		FuncMaps    []template.FuncMap
        // Optional [FuncMapper]
		FuncMapper  FuncMapper
    }
)

// DefaultOptions provides the base options when creating
// a new Renderer.
var DefaultOptions = Options{
	RootDir:     "web/templates",
	GlobPattern: "*.gotmpl",
	FuncMaps:    []template.FuncMap{},
	FuncMapper:  defaultFuncMapper,
}

// Render implements [echo.Renderer].
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	t, err := r.tmpl.Clone()
	if err != nil {
		return err
	}

	funcmap := r.mapper(c)
	_, err = t.Funcs(funcmap).ParseFS(r.dirfs, name)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, name, data)
}

// New creates a new Renderer.
func New(opts *Options) (*Renderer, error) {
	var (
		rootDir  = DefaultOptions.RootDir
		funcMaps = DefaultOptions.FuncMaps
		mapper   = DefaultOptions.FuncMapper
		glob     = DefaultOptions.GlobPattern
	)

	if opts.RootDir != "" {
		rootDir = opts.RootDir
	}
	if opts.GlobPattern != "" {
		glob = opts.GlobPattern
	}
	if opts.FuncMaps != nil {
		funcMaps = opts.FuncMaps
	}
	if opts.FuncMapper != nil {
		mapper = opts.FuncMapper
	}

	dirfs := os.DirFS(rootDir)
	tmpl := template.New("").Funcs(sprig.FuncMap())
	for _, funcmap := range funcMaps {
		_ = tmpl.Funcs(funcmap)
	}

	_, err := tmpl.Funcs(mapper(nil)).ParseFS(dirfs, glob)
	if err != nil {
		return nil, err
	}

	return &Renderer{tmpl, opts.FuncMapper, dirfs}, nil
}

func defaultFuncMapper(echo.Context) template.FuncMap {
	return template.FuncMap{}
}
