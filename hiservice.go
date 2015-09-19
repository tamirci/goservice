package main

import (
	"github.com/tamirci/goservice/Godeps/_workspace/src/github.com/go-martini/martini"
	"log"
)

func main() {
	log.Println("main geldi")
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello protemha"
	})

	m.Run()
}
