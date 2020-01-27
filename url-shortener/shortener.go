package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/lukebrobbs/gophercises/url-shortener/urlShort"
)

func main() {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, http.NewServeMux())
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mapHandler)
}
