// In Go, Map is a data structute that is used to store the
// element in the form of key/value pairs. With the help of map
// we can eaisly modify, retrive and delete value of map using key

package main

import "fmt"

func main() {

	// Creating a map
	myMap := make(map[int]string)

	// Adding key/value pairs in map
	myMap[1] = "One"
	myMap[2] = "Two"
	myMap[3] = "Three"

	fmt.Println("My map : ", myMap)

	fmt.Println("Access the value of a map using key : ", myMap[3])

	// Modify the value of map using key
	myMap[2] = "TwoTwo"
	fmt.Println("Modify the value of map using key : ", myMap)

	// Delete value from map using key
	delete(myMap, 1)
	fmt.Println("Delete value from map using key : ", myMap)

	// Iterate over the map
	for key, value := range myMap {
		fmt.Println("Key : ", key, "Value : ", value)
	}
}
