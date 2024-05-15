package main

import "fmt"

func Find(arr []int, item int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{45, 12, 4, 67, 12, 10}
	item := 9
	fmt.Println(Find(arr, item))
}
