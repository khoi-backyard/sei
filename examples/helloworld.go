package main

import (
	"fmt"
	"log"

	"github.com/zalora/sei"
)

func main() {
	s := sei.New()

	s.Use(func(next sei.HandlerFunc) sei.HandlerFunc {
		return func(c *sei.Context) error {
			fmt.Println("Middlware 1")
			return next(c)
		}
	})

	s.Use(func(next sei.HandlerFunc) sei.HandlerFunc {
		return func(c *sei.Context) error {
			fmt.Println("Middlware 2")
			return next(c)
		}
	})

	s.Use(func(next sei.HandlerFunc) sei.HandlerFunc {
		return func(c *sei.Context) error {
			fmt.Println("Middlware 3")
			return next(c)
		}
	})

	s.GET("/", func(ctx *sei.Context) error {
		fmt.Println("/ handled")
		return ctx.String(200, "Hola")
	})

	s.GET("/json", func(ctx *sei.Context) error {
		return ctx.JSON(200, map[string]interface{}{"ASD": "BDC"})
	})

	log.Fatal(s.Start(":8080"))
}
