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

    SetupViews(e)
    SetupAuth(e)

    e.Logger.Fatal(e.Start(":42069"))
}

