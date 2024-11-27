package main

import (
	"fmt"
	"time"
)

type OperationFn func(int, int)

func main() {

	// Log
	var logAdd, logSubtract func(int, int)
	logAdd = getLogOperation(add)
	logAdd(100, 200)

	logSubtract = getLogOperation(subtract)
	logSubtract(100, 200)

	logMultiply := getLogOperation(func(x, y int) {
		fmt.Println("Multiply result :", x*y)
	})
	logMultiply(100, 200)

	//Profile
	profiledAdd := getProfileOperation(add)
	profiledAdd(100, 200)

	// log & profile
	profiledSubtract := getProfileOperation(getLogOperation(subtract))
	profiledSubtract(100, 200)
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x+y)
}

// ver 2.0

func getLogOperation(oper OperationFn) OperationFn {
	return func(x, y int) {
		fmt.Println("Operation started")
		oper(x, y)
		fmt.Println("Operation completed")
	}
}

func getProfileOperation(oper OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time taken :", elapsed)
	}
}
