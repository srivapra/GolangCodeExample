package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
}

type Dog struct {
	dog string
}
type Cat struct {
	cat string
}

func (d Dog) Sound() string {
	return d.dog

}

func (c Cat) Sound() string {
	return c.cat

}

func Print(a Animal) {
	fmt.Println(a.Sound())
}
func main() {
	d := Dog{"Bhaao"}
	c := Cat{"Meow"}
	Print(d)
	Print(c)
}
