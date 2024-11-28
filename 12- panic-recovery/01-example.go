package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred, err :", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	for range 3 {
		var divisor int
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err != nil {
			fmt.Println("Do not attempt to divide by 0")
			continue
		} else {
			fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

// adapter for 3rd party api (to convert the panic into an error)
func divideAdapter(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient, remainder = x/y, x%y
	return
}
