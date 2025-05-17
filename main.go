package main

import (
	"fmt"
)

func main() {
	demoDefer()
}

func demoDefer(){
	x := 10
	defer fmt.Println("X value: ", x)
	x = 20
	fmt.Println("X changed value: ", x)
}




