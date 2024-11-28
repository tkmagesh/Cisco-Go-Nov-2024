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
	primes := generatePrimes(start, end)
	for _, no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
