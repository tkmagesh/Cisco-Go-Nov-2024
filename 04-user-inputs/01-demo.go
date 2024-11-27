package main

import "fmt"

func main() {

	var userName string
	fmt.Println("Enter the user name :")
	fmt.Scanln(&userName)
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	/*
		var x int
		fmt.Println("Enter a number :")
		if _, err := fmt.Scanln(&x); err != nil {
			fmt.Println("err :", err)
		} else {
			fmt.Println(x)
		}
	*/

	/* var n1, n2 int
	fmt.Println("Enter 2 numbers (separated by space)")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1 + n2) */

}
