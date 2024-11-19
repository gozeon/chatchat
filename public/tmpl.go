package public

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type tmpl struct {
	templates *template.Template
}

func New() *tmpl {
	return &tmpl{
		// templates: template.Must(template.ParseGlob("public/template/*.html")),
		templates: template.Must(template.ParseFS(TemplateDir, "template/*.html")),
	}
}

func (t *tmpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
