package main

import (
	"fmt"
)

func main() {
	names := []string{"Ann", "Boghad"}
	names = append(names, "Vova")
	fmt.Println(names)
	set2Names := []string{"Vova", "Kolya", "Leo"}
	names = append(names, set2Names...)
	fmt.Println(names)

	names = removeElement(names, 2)
	fmt.Println(names)
}

func removeElement(slice []string, index int) []string{
	return append(slice[:index], slice[index+1:]...)
}
