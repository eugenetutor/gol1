package main

import (
	"fmt"
)

func incrementGenerator() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	increment := incrementGenerator()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	anotherIncrement := incrementGenerator()

	fmt.Println(anotherIncrement())
	fmt.Println(anotherIncrement())
	fmt.Println(anotherIncrement())

	double := multiplyBy(2)
	triple := multiplyBy(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

	//example 3
	yourAccount := createBankAccount(1000)
	fmt.Println(yourAccount(500))
	fmt.Println(yourAccount(-200))
}


func multiplyBy(multiplier int) func(int) int {
	 return func(value int) int {
		return value * multiplier
	 }
}

func createBankAccount(initBalance float64) func(float64) float64{
	balance := initBalance
	return func(sum float64) float64{
		balance += sum
		return balance
	}
}


