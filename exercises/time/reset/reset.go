package main

import (
	"log"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2020-03-02T06:57:00.000Z")
	if err != nil {
		log.Fatal(err)
	}
	// Truncate only works with UTC timezone.
	log.Println(t.Truncate(time.Hour * 24))
}
