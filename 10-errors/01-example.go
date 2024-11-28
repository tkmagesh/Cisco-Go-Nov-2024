package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by zero")
	} else {
		fmt.Println("error :", err)
	}
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		// err = errors.New("divisor cannot be zero")
		err = ErrDivideByZero
		return
	}
	quotient, remainder = x/y, x%y
	return
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/
