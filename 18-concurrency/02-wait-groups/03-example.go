package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the counter by 1
	go f1(wg)
	f2()
	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //simulating a time consuming operation
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
