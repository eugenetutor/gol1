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

type Drawable interface{
	Draw()
}

type Resizable interface{
	Resize(factor float64)
}

type GraphicElement interface{
	Drawable
	Resizable
	GetInfo() string
}

type Square struct{
	Size float64
}

func (s *Square) Draw(){
	fmt.Printf("Drawing square with size %.1f \n", s.Size)
}

func (s *Square) Resize(factor float64){
	s.Size *= factor
	fmt.Printf("Drawing square with size %.1f \n", s.Size)
}

func (s Square) GetInfo() string{
	return fmt.Sprintf("Square: size=%.1f", s.Size)
}

func ProcessGraphic(g GraphicElement){
	fmt.Println(g.GetInfo())
	g.Draw()
	g.Resize(1.5)
	g.Draw()
}

func main(){
	square := &Square{Size: 10}
	ProcessGraphic(square)

	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(true)
	PrintAnything([]int{1,2,3})
}


func PrintAnything(value interface{}){
	fmt.Printf("Value: %v, Type: %T \n", value, value)
}


type UserService struct{
	db Database
}