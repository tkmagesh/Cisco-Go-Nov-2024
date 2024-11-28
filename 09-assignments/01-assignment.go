/*
Refactor the logic for generating prime nos to a "generatePrimes" function
Print the generated prime numbers in the main() function
*/

/*
Accept a range from the user and print all the prime numbers between the given range
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
NO_LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue NO_LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
