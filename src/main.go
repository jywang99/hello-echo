package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()
    e.Static("/static", "static")

    // e.GET("/login", )
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", page)
    })
    e.GET("/test", func(c echo.Context) error {
        return c.Render(200, "test", page)
    })
    e.POST("/contacts", getContacts)
    e.DELETE("/contacts/:id", deleteContact)

    e.Logger.Fatal(e.Start(":42069"))
}

