package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	
	http.Handle("/", r)

	srv := http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Response from Home"))
}
