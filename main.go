package main

import (
	"fmt"
)

func main() {
	var value interface{} = 63.3
	switch v := value.(type){
	case string: 
		fmt.Println("String:", v)
	case int: 
		fmt.Println("Int:", v)
	case float64: 
		fmt.Println("Float64:", v)
	default: 
		fmt.Println("Other type")
	}
	
}
