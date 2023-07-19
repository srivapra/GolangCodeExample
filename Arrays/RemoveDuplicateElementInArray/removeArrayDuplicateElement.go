package main

import "fmt"

func removeArrayDuplicateElement(arr []int) []int {

	inResult := make(map[int]bool)

	result := []int{}

	for _, value := range arr {
		if _, ok := inResult[value]; !ok {
			inResult[value] = true
			result = append(result, value)

		}
	}
	return result
}

func main() {

	arr := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	fmt.Println(removeArrayDuplicateElement(arr))

}
