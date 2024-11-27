// Anonymous functions

package main

import (
	"fmt"
	"strings"
)

func main() {

	func() {
		fmt.Println("Hi")
	}()

	func(userName string) {
		fmt.Printf("Hello %q\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	s := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Println(s)

	fn, ln := func(fullName string) (firstName string, lastName string) {
		names := strings.Split(fullName, " ")
		firstName = names[0]
		lastName = names[1]
		return
	}("Magesh Kuppan")
	fmt.Printf("First Name : %q, Last Name : %q\n", fn, ln)
}
