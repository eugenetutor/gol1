package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	currentHour := time.Now().Hour()
	fmt.Println(currentHour)
	if currentHour >= 6 && currentHour < 22{
		fmt.Println("Зараз день")
	}else{
		fmt.Println("Зараз ніч")
	}

	if rand.Intn(2) == 0{
		fmt.Println("Орел")
	}else{
		fmt.Println("решка")
	}

	statusCode := 200
	if statusCode >= 200 && statusCode < 300{
		fmt.Println("Ok")
	}else{
		fmt.Println("no Ok")
	}
}
