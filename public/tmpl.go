package public

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type tmpl struct {
	templates map[string]*template.Template
}

func New() *tmpl {
	templates := make(map[string]*template.Template)

	// 独立页面 、iframe
	templates["chat.html"] = template.Must(template.ParseFS(TemplateDir, "template/chat.html"))
	templates["link.html"] = template.Must(template.ParseFS(TemplateDir, "template/link.html"))
	templates["keyword-default.html"] = template.Must(template.ParseFS(TemplateDir, "template/keyword-default.html"))
	templates["keyword-list.html"] = template.Must(template.ParseFS(TemplateDir, "template/keyword-list.html"))

	// 基础设置
	templates["band.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/band.html"))
	templates["avatar.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/avatar.html"))
	templates["other.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/other.html"))

	// 消息设置
	templates["message.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/message.html"))
	templates["quick.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/quick.html"))
	templates["keyword.html"] = template.Must(template.ParseFS(TemplateDir, "template/_base.html", "template/keyword.html"))

	return &tmpl{
		// templates: template.Must(template.ParseGlob("public/template/*.html")),
		// templates: template.Must(template.ParseFS(TemplateDir, "template/*.html")),
		templates: templates,
	}
}

func (t *tmpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.Execute(w, data)
}
