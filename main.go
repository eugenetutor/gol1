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

type MusicInstrument interface{
	playMusic()
	tuneInstrument()
}

type Piano struct{}

func (p Piano) playMusic(){
	fmt.Println("Дін-дінь-дон")
}

func (p Piano) tuneInstrument(){
	fmt.Println("Налаштовую клавіши")
}

type Guitar struct{}

func (p Guitar) playMusic(){
	fmt.Println("бринь-бринь")
}

func (p Guitar) tuneInstrument(){
	fmt.Println("Підкручую струни")
}

type Drums struct{}

func (p Drums) playMusic(){
	fmt.Println("бум-бум-бум")
}

func (p Drums) tuneInstrument(){
	fmt.Println("підтягую мембрани")
}

func PlaySong(i MusicInstrument){
	fmt.Println("Готується до виступу...")
	i.tuneInstrument()
	fmt.Println("починаємо грати...")
	i.playMusic()
	fmt.Println("браво!")
}

func main(){
	piano := Piano{}
	guitar := Guitar{}
	drums := Drums{}
	PlaySong(piano)
	PlaySong(guitar)
	PlaySong(drums)
}