package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    Templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.Templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        Templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type Count struct {
    Count int
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()
    e.Static("/static", "static")

    count := Count{Count: 0}
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", count)
    })

    e.POST("/count", func(c echo.Context) error {
        count.Count ++
        return c.Render(200, "count", count)
    })

    e.Logger.Fatal(e.Start(":42069"))
}

