package main

import (
	"fmt"
)

func output(input int) {
	fmt.Println("========")
	fmt.Println(&input)
	fmt.Println(input)
}

func main() {
	a := 2
	fmt.Println(&a)
	output(a)
}

// func functionName(par1 type1, par2 type2, ...) return_type{
// 	    //function body
// 	    return value
// }
