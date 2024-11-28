package main

import (
	"fmt"

	/* "github.com/tkmagesh/cisco-go-nov-2024/13-modularity/calculator" */
	"github.com/fatih/color"
	calc "github.com/tkmagesh/cisco-go-nov-2024/13-modularity/calculator"
)

func main() {
	color.Yellow("App started..")
	// fmt.Println("App started..")
	sayHi()
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
	*/
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
}
