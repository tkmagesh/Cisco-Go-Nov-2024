package main

import (
	"fmt"
)

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 2}
	// var productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 2}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 2}
	/*
		productRanks := map[string]int{
			"pen":    5,
			"pencil": 1,
			"marker": 2, // ensure that the last item also has a ","
		}
	*/
	var productRanks = make(map[string]int)
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
	*/
	productRanks["pen"] = 5
	productRanks["pencil"] = 1
	productRanks["marker"] = 2
	fmt.Println(productRanks)

	fmt.Println("iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("check for the existence of a key")
	// keyToCheck := "notepad" // non existent key
	keyToCheck := "pen"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	// remove an item
	// keyToRemove := "pen"
	keyToRemove := "notepad" // non existent key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
