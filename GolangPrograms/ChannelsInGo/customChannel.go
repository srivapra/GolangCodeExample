package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func sendData(ch chan Person, p Person) {

	ch <- p

}

func main() {
	p := Person{Name: "Prashant", Age: 24}

	ch := make(chan Person)

	go sendData(ch, p)

	name := (<-ch).Name
	fmt.Println(name)

}
