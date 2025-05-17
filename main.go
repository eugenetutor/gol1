package main

import (
	"errors"
	"fmt"
)

func main() {

	q,r := divideWithRemainder(13,5)
	fmt.Printf("13 / 5 = %d with remainder %d\n", q, r)
	number, _, _ := getData()
	fmt.Println("Number: ", number)

	_, _, isValid := getData()
	fmt.Println("Correct: ", isValid)

	result, err := divide(10,2)
	if err != nil {
		fmt.Println("Error: ", err)
	}else {
		fmt.Println("Result: ", result)
	}

	result, err = divide(10,0)
	if err != nil {
		fmt.Println("Error: ", err)
	}else {
		fmt.Println("Result: ", result)
	}
}

func divide(a,b float64) (float64, error){
	if b == 0 {
		return 0, errors.New("ділення на нуль")
	}
	return a / b, nil
}

func divideWithRemainder(a,b int) (int, int) {
	q := a/b
	r := a%b
	return q,r
}

func getData() (int, string, bool) {
	return 42, "hello", true
}
