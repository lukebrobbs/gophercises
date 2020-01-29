package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Page is needed to create a html template
type Page struct {
	Title   string
	Body    []string
	Options []option
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type scenario struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type story struct {
	Intro     scenario `json:"intro"`
	NewYork   scenario `json:"new-york"`
	Debate    scenario `json:"debate"`
	SeanKelly scenario `json:"sean-kelly"`
	MarkBates scenario `json:"mark-bates"`
	Denver    scenario `json:"denver"`
	Home      scenario `json:"home"`
}

func loadPage(title scenario) (*Page, error) {

	return &Page{Title: title.Title, Body: title.Story, Options: title.Options}, nil
}

func generateTemplate(s scenario, w http.ResponseWriter) {
	p, err := loadPage(s)
	if err != nil {
		p = &Page{Title: s.Title}
	}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, p)
}

func main() {
	body, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		panic("Unable to read file")
	}
	var d story
	json.Unmarshal(body, &d)

	index := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.Intro, w)
	}
	newYork := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.NewYork, w)
	}
	denver := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.Denver, w)
	}
	seanKelly := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.SeanKelly, w)
	}
	debate := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.Debate, w)
	}
	home := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.Home, w)
	}
	MarkBates := func(w http.ResponseWriter, r *http.Request) {
		generateTemplate(d.MarkBates, w)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/new-york", newYork)
	http.HandleFunc("/denver", denver)
	http.HandleFunc("/sean-kelly", seanKelly)
	http.HandleFunc("/debate", debate)
	http.HandleFunc("/home", home)
	http.HandleFunc("/mark-bates", MarkBates)

	http.ListenAndServe(":9000", nil)
}
