package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for range 20 {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for no := range 20 {
			time.Sleep(500 * time.Millisecond)
			ch <- no
		}
	}()
	return ch
}
