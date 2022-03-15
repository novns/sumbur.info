package main

import (
	"os"
	"sumbur/views/blog"
	"sumbur/views/http_errors"

	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v2"
)

type Config struct {
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

	// Routes

	server.GET("/", blog.BlogGET)

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("Тестовая ошибка")
	})

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
