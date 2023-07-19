package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)

	m[3] = "apple"
	m[9] = "mango"
	m[12] = "date"

	res := make([]int, 0, len(m))
	for k := range m {
		res = append(res, k)

	}

	fmt.Println("Before Sorting ", res)

	sort.SliceStable(res, func(i, j int) bool {
		return m[res[i]] < m[res[j]]
	})

	fmt.Println(res)

}
