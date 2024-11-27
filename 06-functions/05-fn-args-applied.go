package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(i1, i2 int) {
		fmt.Println("Multiply result :", i1*i2)
	}, 100, 200)
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
func logOperation(oper func(int, int), x, y int) {
	fmt.Println("Operation started")
	oper(x, y)
	fmt.Println("Operation completed")
}
