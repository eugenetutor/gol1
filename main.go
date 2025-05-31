package main

import (
	"fmt"
)

//структури - це визначений користувачем тип, який дозволяє групувати, комбінувати елементи різних типів в один тип. 
//struct is value type. NOT REFERENCE TYPE
type Student struct{
	Name string
	Age int 
	Group string
	Avarage float64
}

func main() {
	a := Student{ Name: "Ann", Age: 25}
	b := &a//b- вказівник на а
	b.Name = "Bob"
	fmt.Println(a.Name)
	fmt.Println(b.Name)
}
