package main

import (
	"fmt"
)

func main() {
	// for ініціалізація; умова; лічильник{
	// 	тіло циклу
	// }
	var a int
	for i := 0; i <= 3; i++ {
		a = a + 2
		fmt.Printf("Iteration %d %d \n", i, a)
	}
	fmt.Println(a)
}
