package main

import (
	"github.com/go-martini/martini"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
