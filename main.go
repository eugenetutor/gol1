package main

import (
	"fmt"
)

func main() {
	fmt.Printf("String %d %d \n", 2, 4)
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3,3,4,4,5,5,5,6,6,7,7,7,8,8,8,98,9,8,7,5,4,3,2,4,5,6,7,7,6,5,4,3,3,4,5,56,6,7))
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers{
		total += num
	}
	return total
}
