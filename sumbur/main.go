package main

import (
	"github.com/savsgio/atreugo/v11"
)

func main() {
	server := atreugo.New(atreugo.Config{Addr: "127.0.0.1:8000"})

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.HTTPResponse("sumbur.info")
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
