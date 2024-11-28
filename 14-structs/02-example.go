package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	/* product = id, name, cost */
	/*
		var p Product
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/
	// var p Product = Product{100, "Pen", 10}

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// struct instances are values and hence the assignment creates a copy
	/*
		var p2 Product
		p2 = p
		p2.Cost = 50
		fmt.Printf("p.Cost = %0.2f, p2.Cost = %0.2f\n", p.Cost, p2.Cost)
	*/
	fmt.Println("Before applying 10% discount, p => ", FormatProduct(p))
	ApplyDiscount(&p, 10)
	fmt.Println("After applying 10% discount, p => ", FormatProduct(p))

	newPPtr := new(Product)
	fmt.Println(newPPtr)
}

func FormatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(pPtr *Product, discountPercentage float32) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
