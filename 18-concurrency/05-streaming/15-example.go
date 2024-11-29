package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		for no := range count {
			time.Sleep(500 * time.Millisecond)
			ch <- no
		}
	}()
	return ch
}
