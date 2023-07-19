package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)

}

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++

	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++

	}
	return result

}

func main() {
	arr := []int{23, 41, 12, 91, 1, 3, 5}
	fmt.Println(MergeSort(arr))
}
