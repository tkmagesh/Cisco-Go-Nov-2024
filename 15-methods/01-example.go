package main

import "fmt"

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

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println("Before applying 10% discount, p => ", FormatProduct(p))
	fmt.Println("Before applying 10% discount, p => ", p.Format())

	// ApplyDiscount(&p, 10)
	p.ApplyDiscount(10)

	// fmt.Println("After applying 10% discount, p => ", FormatProduct(p))
	fmt.Println("After applying 10% discount, p => ", p.Format())

}
