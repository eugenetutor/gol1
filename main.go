package main

import (
	"fmt"
)

func main() {
	// for...range
	names := []string{"Alex", "Nikita", "Andrew", "Yevhenii"}

	for index, name := range names{
		fmt.Printf("Index %d: %s\n", index, name)
	}

	for _, name := range names{
		fmt.Printf("Name: %s\n", name)
	}
}
