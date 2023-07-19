package main

import "fmt"

func intersection(a, b []int) []int {
	m := make(map[int]bool)
	for _, x := range a {
		m[x] = true
	}

	var res []int
	for _, x := range b {
		if m[x] {
			delete(m, x)
			res = append(res, x)
		}
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 5, 6}
	b := []int{4, 5, 6, 7, 8, 5}
	fmt.Println(intersection(a, b))
}
