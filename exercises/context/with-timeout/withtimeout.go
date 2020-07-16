package main

import (
	"context"
	"log"
	"time"
)

func doWork(ctx context.Context, name string) {
	for {
		log.Printf("%s doing some work", name)
		time.Sleep(time.Second * 3)
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*8)
	// defer cancelFn()

	go doWork(ctx, "bob")
	go doWork(ctx, "alice")
	go doWork(ctx, "john")

	time.Sleep(time.Second * 12)
}
