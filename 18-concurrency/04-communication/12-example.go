package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		time.Sleep(time.Second)
		fmt.Println(data)
	}()
	ch <- 100
	wg.Wait()
}
