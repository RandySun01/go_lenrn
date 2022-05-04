package main

import "fmt"

type Peo interface {
	Speak(string) string
}

type Stu struct{}

func (stu *Stu) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo Peo = &Stu{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}