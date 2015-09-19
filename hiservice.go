package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("main geldi")
	http.ListenAndServe(":80", nil)
	log.Println("80 geldi")
	http.HandleFunc("/hello", hello)
	log.Println("hello geldi")
}
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello geldi")

	log.Println("hello gecti")
}
