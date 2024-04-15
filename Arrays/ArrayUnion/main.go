package main

import "fmt"

// 2 Pointer Approach
func UnionArray(arr1, arr2 []int) []int {
	i := 0
	j := 0

	res := []int{}

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
			j++
		}

	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++

	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++

	}
	return res

}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	fmt.Println(UnionArray(arr1, arr2))
}
