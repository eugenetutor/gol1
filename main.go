package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 3

	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	dif := a - b
	fmt.Printf("%d - %d = %d\n", a, b, dif)

	prod := a * b
	fmt.Printf("%d * %d = %d\n", a, b, prod)

	quo := a / b
	fmt.Printf("%d / %d = %d\n", a, b, quo)

	rem := a % b
	fmt.Printf("%d %% %d = %d\n", a, b, rem)

	floatQuo := float64(a) / float64(b)
	fmt.Printf("%d / %d = %.1f\n", a, b, floatQuo)

	fmt.Print(math.Pow(float64(a), 2))

}
