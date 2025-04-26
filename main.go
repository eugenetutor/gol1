package main

import "fmt"

func main(){

	value := 10
	if true {
		value := 20 //нова змінна
		fmt.Println(value)
	}

	fmt.Println(value)


}