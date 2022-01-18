package main

import (
	"fmt"
	"net/http"
)

func contanctHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "email me at <a href=\"mailto:test@gmail.com\">test@gmail.com</a>.")
}

func handlerFunc(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "page not found", http.StatusNotFound)
}

type CustomRouter struct{}

func (router CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "page not found", http.StatusNotFound)
}

func main() {
	var my_router CustomRouter
	fmt.Println("starting the server on :3000...")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", my_router)
}
