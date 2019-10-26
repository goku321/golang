package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHanlder struct {
	mu sync.Mutex
	n  int
}

func (h *countHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.n++
	fmt.Fprintf(w, "Count is %d\n", h.n)
}

func main() {
	handler := countHanlder{}
	http.Handle("/count", &handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
