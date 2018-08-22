package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

var _ Animal = new(Dog)
var _ Animal = &Dog{}
var _ Animal = new(Cat)
//var _ Animal = Cat{} //Cat does not implement Animal (Speak method has pointer receiver)

func main() {
	var dog *Dog

	animals := []Animal{dog, Cat{}, new(Llama), JavaProgrammer{}} //Cat does not implement Animal (Speak method has pointer receiver)
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	animals2 :=[]Animal2{ &ScalaProgrammer{}}
	for _, animal := range animals2 {
		fmt.Println(animal.Speak())
		fmt.Println(animal.SayHi())
	}
}



type Animal2 interface {
	Speak() string
	SayHi() string
}

type ScalaProgrammer struct {
}

func (s *ScalaProgrammer) Speak() string{ // Use pointer receiver
	return "scala aswone"
}

func (s ScalaProgrammer) SayHi() string{ // Use value receiver
	return "hi scala"
}