package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1 through the builtin scheduler
	f2()
	// block the execution of main() so that the scheduler can look for other goroutines scheduled and execute them (cooperative multitasking)
	// DO NOT DO THIS in your applications (poor man's synchronization techniques)
	time.Sleep(time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second) //simulating a time consuming operation
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
