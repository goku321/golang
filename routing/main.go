package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	// Add middleware
	r.Use(loggingMiddleware)
	
	http.Handle("/", r)

	srv := http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}

// HomeHandler - root handler
func HomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Response from Home"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
