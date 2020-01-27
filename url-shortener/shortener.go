package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/redirect" {
		http.Redirect(w, r, "/hello", 301)
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
