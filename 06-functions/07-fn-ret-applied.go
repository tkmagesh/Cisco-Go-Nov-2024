package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(i1, i2 int) {
			fmt.Println("Multiply result :", i1*i2)
		}, 100, 200)
	*/

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
/*
Operation started
Add Result : 300
Operation completed

Operation started
Subtract Result : -100
Operation completed
*/

/*
func logAdd(x, y int) {
	fmt.Println("Operation started")
	add(x, y)
	fmt.Println("Operation completed")
}

func logSubtract(x, y int) {
	fmt.Println("Operation started")
	subtract(x, y)
	fmt.Println("Operation completed")
}
*/

// applying commanility-variability
/*
func logOperation(oper func(int, int), x, y int) {
	fmt.Println("Operation started")
	oper(x, y)
	fmt.Println("Operation completed")
}
*/

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Operation started")
		oper(x, y)
		fmt.Println("Operation completed")
	}
}

func getProfileOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time taken :", elapsed)
	}
}
