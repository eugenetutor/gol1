package main

import "fmt"

func main(){
	const Pi = 3.14159
	const MaxUsers = 1000

	const (
		StatusOK = 200
		StatusNotFound = 404
		StatusInternalErro = 500
	)

	StatusOK = 300//error
	fmt.Println(StatusOK)
}