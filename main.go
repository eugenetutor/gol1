package main

import (
	"fmt"
)

func main() {
	//слайси - це динамічний масив, який може змінювати свій розмір під час виконання програми
	//var імʼя []тип
	// var numbers []int

	arr := [5]int{1,2,3,4,5}

	slice1 := arr[1:4]
	fmt.Println(slice1)

	slice2 := arr[:3]
	slice3 := arr[2:]
	slice4 := arr[:]
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
