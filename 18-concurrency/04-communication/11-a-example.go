package main

import (
	"fmt"
	"sync"
)

// share memory by communicating

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
