/* struct compostion */
package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
*/

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (pPtr *Product) ApplyDiscount(discountPercentage float32) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
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

// overriding Format() method
/*
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}
*/

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func main() {

	var milk = NewPerishableProduct(200, "Milk", 20, "2 Days")

	// fmt.Println(milk.Format())
	fmt.Println(milk)
	milk.ApplyDiscount(10)
	// fmt.Println(milk.Format())
	fmt.Println(milk)

}
