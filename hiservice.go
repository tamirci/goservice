package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("main geldi")
	http.ListenAndServe(":80", nil)
	log.Println("80 geldi")

	log.Println("hello geldi")
}
