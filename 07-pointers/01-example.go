package main

import "fmt"

func main() {
	var x int
	x = 100

	var xPtr *int
	xPtr = &x //pointer to x
	fmt.Println("xPtr : ", xPtr)

	// dereferencing - accessing underlying the data using the address
	fmt.Println("*xPtr : ", *xPtr)

	*xPtr += 200
	fmt.Println("x : ", x)

	fmt.Println("Before incrementing, x :", x)
	increment(&x)
	fmt.Println("After incrementing, x :", x)

	a, b := 100, 200
	fmt.Printf("Before swapping, a = %d & b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping, a = %d & b = %d\n", a, b)
}

func increment(no *int) {
	fmt.Println("&no :", no)
	*no++
}

func swap(n1, n2 *int) { //NO return result
	*n1, *n2 = *n2, *n1
}
