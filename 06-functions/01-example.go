package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	s := getGreet("Suresh", "Kannan")
	fmt.Println(s)

	fn, ln := getNames("Magesh Kuppan")
	fmt.Printf("First Name : %q, Last Name : %q\n", fn, ln)
}

func sayHi() {
	fmt.Println("Hi")
}

func sayHello(userName string) {
	fmt.Printf("Hello %q\n", userName)
}

/*
	func greet(firstName string, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
*/
func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreet(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
}

/*
func getNames(fullName string) (string, string) {
	var firstName, lastName string
	names := strings.Split(fullName, " ")
	firstName = names[0]
	lastName = names[1]
	return firstName, lastName
}
*/
// Named results
func getNames(fullName string) (firstName string, lastName string) {
	names := strings.Split(fullName, " ")
	firstName = names[0]
	lastName = names[1]
	return
}
