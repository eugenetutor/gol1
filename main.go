package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "Привіт"
	fmt.Println(len(text))
	for i:=0; i< len(text); i++{
		fmt.Printf("Байт %d: %d\n", i, text[i])
	}

	fmt.Println(utf8.RuneCountInString(text))

	//rune - псевдонім для типу int32.Unicode символи

	var runa2 rune = '人'
	fmt.Printf("тип:%T, значення:%v\n", runa2, runa2)

	var symbol rune = '世'
	fmt.Printf("Symbol: %c\n", symbol)
	fmt.Printf("Symbol: %d\n", symbol)
	fmt.Printf("Symbol: %U\n", symbol)
	fmt.Printf("Symbol: %T\n", symbol)

	var symbol2 rune = '🗿'
	fmt.Printf("Symbol: %c\n", symbol2)
	fmt.Printf("Symbol: %d\n", symbol2)
	fmt.Printf("Symbol: %U\n", symbol2)
	fmt.Printf("Symbol: %T\n", symbol2)
}
