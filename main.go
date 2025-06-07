package main

import (
	"fmt"
)

//Інтерфейси - це колекція сигнатур  методів. Це контракт, який визначає, які методи повинен мати тип, але НЕ РЕАЛІЗУЄ ЇХ
// type НазваІнтерфейсу interface{
// 		НазваМетоду(параметр) повертає
// 		ІншийМетоду(параметр) повертає
// }

//Duck typing - "якщо щось ходить як качка і крякає як качка, то це качка" 

type Animal interface{
	MakeSound() string
}

type Dog struct{
	Name string
	Breed string
}

type Bird struct{
	Name string
	Species string
}

func (d Dog) MakeSound() string{
	return "woof!"
}

func (d Bird) MakeSound() string{
	return "tweet!"
}

// func PrintDogSound(d Dog){
// 	fmt.Println(d.MakeSound())
// }

// func PrintBirdSound(b Bird){
// 	fmt.Println(b.MakeSound())
// }

func printAnimalSound(a Animal){
	fmt.Println(a.MakeSound())
}

func main() {
	bird := Bird{Name: "Kesha", Species: "Parrot"}
	dog := Dog{Name: "Boris", Breed: "Golden Retruever"}

	// PrintBirdSound(bird)
	// PrintDogSound(dog)

	printAnimalSound(bird)
	printAnimalSound(dog)
}