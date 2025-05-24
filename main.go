package main

import (
	"fmt"
)

func main() {
	//масиви - це нумерований послідовність  елементів ОДНОГО ТИПУ з ФІКСОВАНОЮ довжиною
	//var імʼя [розмір]тип
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(&numbers[0])
	fmt.Println(&numbers[1])
	fmt.Println(&numbers[2])
	fmt.Println(&numbers[3])
	fmt.Println(&numbers[4])

	var fruits [4]string = [4]string{"apple", "orange", "banana", "peach"}
	fmt.Println(&fruits[0])
	fmt.Println(&fruits[1])
	fmt.Println(&fruits[2])
	fmt.Println(&fruits[3])

	names := [...]string{"Ann", "Bogdan", "Vika"}
	names[2] = "Kleo"
	fmt.Println(names)
	fmt.Println(len(names))
}
