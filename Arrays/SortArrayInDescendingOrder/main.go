package main

import (
	"fmt"
	"sort"
)

func SortInDescendingOrder(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}
func main() {
	arr := []int{4, 1, 8, 10, 40, 5, 12}
	fmt.Print(SortInDescendingOrder(arr))

}
