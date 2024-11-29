package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- add(100, 200)
	}()
	result := <-ch
	fmt.Println(result)
}

// 3rd party api
func add(x, y int) int {
	return x + y
}
