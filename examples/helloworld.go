package main

import (
	"log"

	"github.com/zalora/sei"
)

func main() {
	s := sei.New()

	s.GET("/", func(ctx *sei.Context) error {
		return ctx.String(200, "Hola")
	})

	s.GET("/json", func(ctx *sei.Context) error {
		return ctx.JSON(200, map[string]interface{}{"ASD": "BDC"})
	})

	log.Fatal(s.Start(":8080"))
}
