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
	var x int = 42
	var p *int = &x
	fmt.Println(*p)
	*p = 20
	fmt.Println(*p)

	x, y := 10, 20
	swap(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x,y)
}

func swap(a,b *int){
	*a, *b = *b, *a
}
