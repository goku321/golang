package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing17/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener: started: listening on :4000")
	http.ListenAndServe(":4000", nil)
}