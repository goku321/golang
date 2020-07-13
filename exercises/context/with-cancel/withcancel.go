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
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	go doWork(ctx, "bob")
	go doWork(ctx, "joe")
	go doWork(ctx, "alice")

	time.Sleep(time.Second * 5)
}
