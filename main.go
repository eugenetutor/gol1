package main

import (
	"fmt"
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
}
