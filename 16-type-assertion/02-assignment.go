package main

/*
Hint : use strconv.Atoi() to convert integers in string format to int
*/
import "fmt"

func main() {
	fmt.Println(sum(10))                      // => 10
	fmt.Println(sum(10, "20"))                //=> 30
	fmt.Println(sum(10, 20, "30", 40, "abc")) //=> 100
	fmt.Println(sum())                        //=>0
}

func sum(nos ...int) int {
	var result int
	for _, no := range nos {
		result += no
	}
	return result
}
