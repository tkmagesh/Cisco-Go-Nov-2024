package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	var count int
	flag.IntVar(&count, "count", 0, "Number of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("About to spin up %d goroutines.. Hit ENTER to start!\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1)
		go fn(wg, id)
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second) //simulating a time
	fmt.Printf("fn[%d] completed\n", id)
}
