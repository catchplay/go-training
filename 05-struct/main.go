package main

import (
	"fmt"

	"github.com/catchplay/go-training/05-struct/person"
)

func main() {
	fmt.Println(person.New("Bob", 20))

	p := person.New("Eason", 32)
	p.SetName("Ethan")
	p.SetAge(10)

	fmt.Println(p)
}
