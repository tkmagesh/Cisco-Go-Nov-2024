// functions as values to variables

package main

import (
	"fmt"
	"strings"
)

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello %q\n", userName)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	var getGreet func(string, string) string
	getGreet = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}
	s := getGreet("Suresh", "Kannan")
	fmt.Println(s)

	var getNames func(string) (string, string)
	getNames = func(fullName string) (firstName, lastName string) {
		names := strings.Split(fullName, " ")
		firstName = names[0]
		lastName = names[1]
		return
	}
	fn, ln := getNames("Magesh Kuppan")
	fmt.Printf("First Name : %q, Last Name : %q\n", fn, ln)

	var f func()
	f = f1
	f()
}

func f1() {
	fmt.Println("f1 invoked")
}
