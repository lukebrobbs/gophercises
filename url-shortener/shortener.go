package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/lukebrobbs/gophercises/url-shortener/urlShort"
)

func logger(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(" %v called\n", r.RequestURI)
		h.ServeHTTP(w, r)
	}
}

func main() {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	mapHandler := urlshort.MapHandler(pathsToUrls, http.NewServeMux())
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", logger(yamlHandler))
}
