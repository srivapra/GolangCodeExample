package main

import "fmt"

// Animal is a base struct representing an animal.
type Animal struct {
	Name string
	Age  int
}

// Eat is a method defined on the Animal struct.
func (a Animal) Eat() {
	fmt.Println("The animal is eating.")
}

// Dog is a struct that embeds the Animal struct.
type Dog struct {
	Animal
	Breed string
}

// Bark is a method defined on the Dog struct.
func (d Dog) Bark() {
	fmt.Println("The dog is barking.")
}

func main() {
	// Creating an instance of Dog.
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Labrador",
	}

	// Accessing fields and methods of the embedded Animal struct.
	fmt.Println("Name:", dog.Name)
	fmt.Println("Age:", dog.Age)
	dog.Eat()

	// Calling methods specific to the Dog struct.
	dog.Bark()
}
