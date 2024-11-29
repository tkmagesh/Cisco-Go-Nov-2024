package main

/*
Hint : use strconv.Atoi() to convert integers in string format to int
*/
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))                      // => 10
	fmt.Println(sum(10, "20"))                //=> 30
	fmt.Println(sum(10, 20, "30", 40, "abc")) //=> 100
	fmt.Println(sum())                        //=>0
}

func sum(list ...interface{}) int {
	var result int
	for _, item := range list {
		switch val := item.(type) {
		case int:
			result += val
		case string:
			if no, err := strconv.Atoi(val); err == nil {
				result += no
			}
		}
	}
	return result
}
