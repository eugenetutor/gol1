package main

import (
	"fmt"
	"time"
)

func main() {
	// switch вираз {
	// case знач1:
	// 		//код1
	// case знач2:
	// 		//код2
	// case знач3:
	// 		//код3
	// default:
	// 		//код, якщо нічого не підішло
	// }

	rating := 4
	switch rating{
	case 5:
		fmt.Println("Дякуємо щиро за ваг відгук!")
	case 4:
		fmt.Println("Дякуємо за високу оцінку!")
	case 3:
		fmt.Println("Дякуємо! Ми працюємо над покращенням!")
	case 1,2:
		fmt.Println("Нам шкода, що ви залишилися незадоволеними")
	default:
		fmt.Println("Некоректна оцінка")
	}
}
