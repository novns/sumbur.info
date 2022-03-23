package main

import (
	"os"
	"sumbur/views/auth"
	"sumbur/views/blog"
	"sumbur/views/http_errors"
	"sumbur/views/static"

	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth   auth.Config
	Server atreugo.Config
}

func main() {
	// Configuration

	config := Config{
		Server: atreugo.Config{
			NoDefaultContentType: true,

			MethodNotAllowedView: http_errors.NotFoundView,
			NotFoundView:         http_errors.NotFoundView,
			PanicView:            http_errors.PanicView,
		},
	}

	file, err := os.ReadFile(os.Getenv("SUMBUR_CONFIG"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	// Server

	server := atreugo.New(config.Server)

	auth.Init(&config.Auth)
	server.UseBefore(auth.Check)

	static.Debug = config.Server.Debug

	// Routes

	server.GET("/", blog.BlogGET)

	server.POST("/auth", auth.AuthPOST)

	server.GET("/restrict", auth.RestrictedGET).
		UseBefore(auth.Restrict)

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("Тестовая ошибка")
	})

	server.Static("/static", "static")
	server.GET("/static-stamp/{stamp}/{path:*}", static.StaticStampGET)

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
