package models

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

type PerishableProduct struct {
	Product
	Expiry   string
	category string // private member
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry:   expiry,
		category: "Food",
	}
}

// overriding Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q, Category = %q", pp.Product.Format(), pp.Expiry, pp.category)
}

// private to "models" package
func (pp PerishableProduct) whoAmI() string {
	return "I am a Perishable Product"
}
