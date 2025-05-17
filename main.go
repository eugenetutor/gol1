package main

import (
	"fmt"
)

func main() {
	rectangleArea := calcArea(5.0, 3.0)
	fmt.Println("Rectangle Area: ", rectangleArea)

	checkValue(5)
	demoShadowing()
}

func calcArea(w, h float64) float64 {
	area := w * h
	return area
}

func checkValue(n int) {
	if n > 0 {
		message := "Lion"
		fmt.Println(message)
	}
	// fmt.Println(message)
}

func demoShadowing(){
	x := 10
	fmt.Println("initial x: ",x)
	if true{
		x := 20
		fmt.Println(&x)
		fmt.Println("x in if block: ",x)
	}
	fmt.Println("x after if block: ",x)
	fmt.Println(&x)
}

