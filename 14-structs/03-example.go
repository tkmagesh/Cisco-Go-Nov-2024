/* struct compostion */
package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	// var milk = PerishableProduct{Product{200, "Milk", 20}, "2 Days"}
	/*
		var milk = PerishableProduct{
			Product: Product{
				Id:   200,
				Name: "Milk",
				Cost: 20,
			},
			Expiry: "2 Days",
		}
	*/
	var milk = NewPerishableProduct(200, "Milk", 20, "2 Days")
	fmt.Println(milk)

	// Accessing the members
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)
}

func FormatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(pPtr *Product, discountPercentage float32) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
