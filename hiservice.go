package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("main geldi")
	http.HandleFunc("/hello", hello)
	log.Println("hello geldi")
	http.ListenAndServe(":8080", nil)
	log.Println("80 geldi")
}
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello geldi")
	fmt.Fprintf(w, "hi!")
	log.Println("hello gecti")
}
