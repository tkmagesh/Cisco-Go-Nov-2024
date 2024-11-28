package main

import (
	"errors"
	"fmt"
)

var ErrF1 = errors.New("f1 error")
var ErrF2 = errors.New("f2 error")

func main() {
	err := f1()
	fmt.Println("[main] error :", err)
	if errors.Is(err, ErrF1) {
		fmt.Println("f1 resulted in an error")
	}
	if errors.Is(err, ErrF2) {
		fmt.Println("f2 resulted in an error")
	}
}

func f1() error {
	err := f2()
	return fmt.Errorf("%w %w", ErrF1, err)

	// return ErrF1
}

func f2() error {
	return ErrF2
}
