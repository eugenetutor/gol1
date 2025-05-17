package main

import (
	"fmt"
)

func main() {
	x := 20
	func(){
		x = 25
		fmt.Println("Hello from anonymous func", x)
	}()

	func(name string){
		fmt.Println("Hello from anonymous func", name)
	}("Golang")

	square := func(n int) int {
		return n * n
	}

	fmt.Println(square(8))
}
