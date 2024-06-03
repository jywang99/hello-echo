package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
)

type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port int `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Host string `yaml:"host"`
        Port int `yaml:"port"`
    } `yaml:"db"`
}

func ReadConfig() Config {
    f, err := os.Open("conf/config.yml")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var cfg Config
    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(&cfg)
    if err != nil {
        log.Fatal(err)
    }

    return cfg
}
var config = ReadConfig()

func SetupGetConfig(e *echo.Echo) {
    e.GET("/config", func(c echo.Context) error {
        return c.JSONPretty(200, config, "  ")
    })
}

