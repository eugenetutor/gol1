package main

import (
	"fmt"
)

func main() {
	// for...range
	marks := map[string]int{
		"Alex": 20,
		"Nikita": 50,
		"Andrew": 30,
		"Yevhenii": 2,
	}

	for name, mark := range marks{
		fmt.Printf("%s: %d\n", name, mark)
	}
	
	for name := range marks{
		fmt.Println("Name: ", name)
		if name == "Alex"{
			break
		}
	}
	
	outerLoop:
	for i := 0; i<3; i++{
		for j:=0; j<3; j++{
			if i == 2 && j == 2{
				fmt.Println("Переривання зовнішнього циклу")
				break outerLoop
			}
			fmt.Printf("i: %d, j: %d\n", i,j)
		}
	}

	for i := 0; i<10; i++{
		if i % 2 != 0{
			continue
		}
		fmt.Println("Парне число", i)
	}

	for i := 2; i < 10; i++ {
		
		if 21%i == 0 {
			fmt.Println(i)
		}
		if i == 9 {
			fmt.Println(i)
		}
	}

}
