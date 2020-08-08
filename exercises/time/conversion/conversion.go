package main

import (
	"log"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2012-12-12T15:04:05Z")
	if err != nil {
		log.Fatal(err)
	}
	ft := t.Format("2006-01-02 15:04:05")
	log.Println(ft)

	log.Println(time.Now())

	t, err = time.Parse("Mon Jan 02 2006 15:04:05 MST-0700", "Tue Jun 11 2019 13:26:45 GMT+00+0000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t)
}
