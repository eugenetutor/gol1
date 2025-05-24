package main

import (
	"fmt"
)

func main() {
	names := make([]string, 3, 5)// ["", "", "", "",""]
	names[0] = "Ann"
	names[1] = "Vova"
	names[2] = "Oleg"
	fmt.Println(names)
	fmt.Println("Length", len(names))
	fmt.Println("Capacity", cap(names))

	s := make([]int, 0, 3)
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))

	for i := 1; i <= 7; i++ {
		s = append(s, i)
		fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))
	}
}

