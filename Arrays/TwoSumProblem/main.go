package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for arrayIndex, arrayValue := range nums {
		if mapKey, mapValue := m[target-arrayValue]; mapValue {
			return []int{mapKey, arrayIndex}
		}
		m[arrayValue] = arrayIndex
	}
	return nil

}

func main() {

	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))

}
