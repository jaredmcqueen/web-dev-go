package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func contanctHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "email me at <a href=\"mailto:test@gmail.com\">test@gmail.com</a>.")
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h3>this is a webpage</h3>")
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")
	slugParam := chi.URLParam(r, "slug")

	fmt.Fprintf(w, "date was %v, slug was %v", dateParam, slugParam)
}

func main() {
	r := chi.NewRouter()
	r.Get("/contact", contanctHandler)
	r.Get("/", homeHandler)
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	r.Get("/articles/{date}-{slug}", getArticle)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
