package main

import "fmt"

func main() {
	// var nos [5]int // memory is allocated for 5 ints and initialized with the 'zero values' of int (0)
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	// nos := [5]int{3, 1, 4}
	fmt.Println(nos)

	// iterating an array using indexer
	fmt.Println("iterating an array using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating an array using range
	fmt.Println("iterating an array using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		nos2 := nos // creating a copy of nos
		nos2[0] = 9999
		fmt.Printf("nos2[0] = %d, nos[0] = %d\n", nos2[0], nos[0])
	*/

	nos2 := &nos
	nos2[0] = 9999
	fmt.Printf("nos2[0] = %d, nos[0] = %d\n", nos2[0], nos[0])

	sort(&nos)
	fmt.Println(nos)
}

func sort(list *[5]int) { // NO return results
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
