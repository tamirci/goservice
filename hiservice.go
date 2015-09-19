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
	if r.URL.Path != "/hello" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	log.Println("hello geldi")
	fmt.Fprintf(w, "hi!")
	log.Println("hello gecti")
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
