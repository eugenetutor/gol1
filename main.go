package main

import (
	"fmt"
)

func main() {
	//мапа - це невпорядкована колекція пар ключ-значень, де ключі УНІКАЛЬНІ

	// var name map[KeyType]ValueType
	// var scores map[string]int

	userAges := map[string]int{
		"Ann": 25,
		"Bohdan": 30,
		"Vika": 22,
	}
	fmt.Println(userAges)
	contacts := make(map[string]string)

	contacts["Mike"] = "+380505005050"
	contacts["Natalie"] = "+380666006060"
	fmt.Println(contacts)

	mikePhone, exists := contacts["Leo"]
	if exists {
		fmt.Println(mikePhone)
	} else {
		fmt.Println("nothing")
	}
	
	delete(contacts, "Mike")
	fmt.Println(contacts)
}

