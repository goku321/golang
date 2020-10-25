package main

import (
	"fmt"
	"log"
	"time"
)

func slowOperation() {
	defer trace("slowOperation")()
	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func main() {
	_ = double(4)
}
