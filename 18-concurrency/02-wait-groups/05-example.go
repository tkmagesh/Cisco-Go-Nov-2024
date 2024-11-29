package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go f1(wg)
	wg.Wait() // result in a deadlock
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //simulating a time consuming operation
	fmt.Println("f1 completed")
}
