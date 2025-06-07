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

type MessageSender interface{
	Send(message string)
}

type EmailSender struct{}
func (e *EmailSender) Send(message string){
	fmt.Printf("EMail: %s \n", message)
}

type SMSSender struct{}
func (e *SMSSender) Send(message string){
	fmt.Printf("SMS: %s \n", message)
}

type TelegramSender struct{}
func (e *TelegramSender) Send(message string){
	fmt.Printf("Telegram: %s \n", message)
}

type NotificationService struct{
	sender MessageSender
}

func (n *NotificationService) ChangeSender(newSender  MessageSender){
	n.sender = newSender
}

func (n *NotificationService) SendNotification(msg string){
	n.sender.Send(msg)
}

func main(){
	email := &EmailSender{}
	sms := &SMSSender{}
	telegram := &TelegramSender{}

	service := &NotificationService{sender: email}
	service.SendNotification("привіт чеерез email")

	service.ChangeSender(sms)
	service.SendNotification("привіт чеерез sms")

	service.ChangeSender(telegram)
	service.SendNotification("привіт чеерез telegram")
}


