package main

import (
	"net/http"
	"html/templates"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)
	}
	files := []string{"templates/layout.html",
					  "templates/navbar.html",
					  "templates/index.html",}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}