package main

import "fmt"

func Add(add func(x, y int) int) {
	fmt.Println("Additiion : ", add(5, 6))
}

func main() {
	result := func(x, y int) int {
		res := x + y
		return res
	}

	Add(result)

}
