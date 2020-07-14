package main

import (
	"log"
	"sync"
	"time"
)

type pool struct {
	work chan int
	wg   sync.WaitGroup
}

func new(concurrency int) *pool {
	p := &pool{
		work: make(chan int),
	}

	p.wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for w := range p.work {
				doWork(w)
			}
			p.wg.Done()
		}()
	}

	return p
}

func (p *pool) run(id int) {
	p.work <- id
}

func (p *pool) shutdown() {
	close(p.work)
	p.wg.Wait()
}

func doWork(id int) {
	log.Printf("%d doing some imaginary work", id)
	time.Sleep(time.Second * time.Duration(id))
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := new(2)

	for _, t := range tasks {
		p.run(t)
	}

	time.Sleep(time.Second * 10)
}
