package main

import "fmt"

func main() {
	/* product = id, name, cost */
	/*
		var p struct {
			Id   int
			Name string
			Cost float32
		}
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/
	var p struct {
		Id   int
		Name string
		Cost float32
	} = struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// struct instances are values and hence the assignment creates a copy
	/*
		var p2 struct {
			Id   int
			Name string
			Cost float32
		}
		p2 = p
		p2.Cost = 50
		fmt.Printf("p.Cost = %0.2f, p2.Cost = %0.2f\n", p.Cost, p2.Cost)
	*/
	fmt.Println("Before applying 10% discount, p => ", FormatProduct(p))
	ApplyDiscount(&p, 10)
	fmt.Println("After applying 10% discount, p => ", FormatProduct(p))
}

func FormatProduct(p struct {
	Id   int
	Name string
	Cost float32
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(pPtr *struct {
	Id   int
	Name string
	Cost float32
}, discountPercentage float32) {
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
