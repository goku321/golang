package main

import (
	"fmt"
	"regexp"
	"time"

	"log"
)

func main() {
	l := time.FixedZone("GMT", 14400)
	fmt.Println(l.String())

	input := "Tue Jun 11 2019 13:26:45 GMT+04:00"
	re := regexp.MustCompile(`[A-Z]+[+-0-9]+.*`)
	matches := re.FindStringSubmatch(input)
	fmt.Println("Found timezone string: ", matches)
	
	loc, err := time.LoadLocation(matches[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(loc)
}
