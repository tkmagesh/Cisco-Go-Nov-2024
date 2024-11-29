package main

import (
	"fmt"

	/* "github.com/tkmagesh/cisco-go-nov-2024/13-modularity/calculator" */
	"github.com/fatih/color"
	calc "github.com/tkmagesh/cisco-go-nov-2024/13-modularity/calculator"
	"github.com/tkmagesh/cisco-go-nov-2024/13-modularity/models"
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

	var milk = models.NewPerishableProduct(200, "Milk", 20, "2 Days")

	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())

	// fmt.Println(milk.category) // private member "category" is not accessible outside the "models" package

}
