package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Renderer = newTemplate()
    e.Static("/static", "static")

    SetupViews(e)
    SetupAuth(e)
    SetupGetConfig(e)
    SetupUserQuery(e)

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Server.Port)))
}

