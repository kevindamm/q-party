package service

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

func NewRenderer() echo.Renderer {
	return TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.ht_")),
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (site TemplateRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return site.templates.ExecuteTemplate(w, name, data)
}
