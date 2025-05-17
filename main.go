package main

import (
	"fmt"
)

func main() {
	demoDefer()
}

func demoDefer(){
	//LIFO - (Last in, First out)
	defer fmt.Println("Last output")
	defer fmt.Println("Second output")

	fmt.Println("First output")
}




