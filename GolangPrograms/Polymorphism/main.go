package main

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
	d string
}

type Cat struct {
	c string
}

func (dog Dog) Bark() {
	fmt.Println("Dog is barking")
}

func (cat Cat) Bark() {
	fmt.Println("Cat is barking")
}

func Print(i Animal) {
	i.Bark()
}

func main() {
	d := Dog{"Labra"}
	c := Cat{"Cat"}
	Print(d)
	Print(c)
}
