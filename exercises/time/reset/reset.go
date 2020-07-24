package main

import (
	"log"
	"time"
)

func resetTime(t time.Time) time.Time {
	truncated := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return truncated
}

func main() {
	t, err := time.Parse(time.RFC3339, "2020-03-02T06:57:00.000Z")
	if err != nil {
		log.Fatal(err)
	}
	// Truncate only works with UTC timezone.
	log.Println(t.Truncate(time.Hour * 24))

	// Let's try for a different timezone.
	t = time.Now()
	// Will it work?
	log.Println(t.Truncate(time.Hour * 24))

	// Let's try to with our custom function.
	log.Println(resetTime(t))
}
