package main

import (
	"fmt"
	"math/rand"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (pPtr *Product) ApplyDiscount(discountPercentage float32) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	// var x interface{}
	var x any
	// x = 100
	// x = "Eiusmod excepteur sunt ea dolor do Lorem eiusmod sint amet exercitation."
	// x = true
	// x = 99.99
	// x = Product{Id: 100, Name: "Pen", Cost: 10}
	// fmt.Println(x)

	// x = getValue()
	x = 200
	// x = "Labore sunt et eu nisi et nulla enim ex."
	// y := x.(int) * 2

	// type assertion using "if"
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	fmt.Println("type assertion using \"type switch\"")
	// x = 100
	// x = "Tempor nostrud Lorem nostrud nisi aute laboris duis ut aute."
	// x = true
	// x = 99.99
	// x = Product{Id: 100, Name: "Pen", Cost: 10}
	x = struct{}{}
	// type assertion using "type switch"
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.09 =", val*0.09)
	case Product:
		fmt.Println("x is a product, x.Format() =", val.Format())
	default:
		fmt.Println("x is an unknown type")
	}

}

func getValue() interface{} {
	if rand.Intn(20)%2 == 0 {
		return 100
	}
	return "Esse mollit nulla ad mollit elit nulla tempor."
}
