package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating a slice using indexer
	fmt.Println("iterating a slice using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating a slice using range
	fmt.Println("iterating a slice using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// adding new items
	nos = append(nos, 90)
	fmt.Println(nos)

	nos = append(nos, 50, 20, 40, 30, 10)
	fmt.Println(nos)

	hundreds := []int{300, 100, 200}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos2[0] = %d, nos[0] = %d\n", nos2[0], nos[0])

	sort(nos)
	fmt.Println("len(nos) : ", len(nos))
	fmt.Println(nos)

	// slicing
	fmt.Println("slicing")
	nosSlice := nos[3:6]
	nosSlice[0] = 8888
	fmt.Println("nos[3:6] =", nos[3:6])
	fmt.Println("nos[3:] =", nos[3:])
	fmt.Println("nos[:6] =", nos[:6])
}

func sort(list []int) { // NO return results
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
