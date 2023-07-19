// Slice is similar to an array. Array has fixed size, but slice doesn't have
// fixed size, We can grow and shrink the slices

package main

import "fmt"

func main() {

	var myArray = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Array : ", myArray)

	// Creating a slice from an array
	mySlice := myArray[3:6]
	fmt.Println("Creating slice from an array : ", mySlice)

	// Creating a slice without an array
	slice := []string{"This", "is", "a", "string"}
	fmt.Println("Creating a slice without an array : ", slice)

	// Add more values in slice
	slice = append(slice, "after", "modification")
	fmt.Println("Grow slice : ", slice)

	slice = append(slice[2:4])
	fmt.Println("Shrink slice : ", slice)

	// Creating a slice using make
	sliceUsingMake := make([]int, 3)
	sliceUsingMake[0] = 1
	sliceUsingMake[1] = 2
	sliceUsingMake[2] = 3

	sliceUsingMake = append(sliceUsingMake, 4, 5)

	fmt.Println("Slice using make : ", sliceUsingMake)

	// Remove value from slice based on index
	index := 2

	sliceUsingMake = append(sliceUsingMake[:index], sliceUsingMake[index+1:]...)
	fmt.Println("Remove values from slice base on index : ", sliceUsingMake)

}
