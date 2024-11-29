package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		for no := range count {
			ch <- no
		}
		close(ch)
	}()
	return ch
}
