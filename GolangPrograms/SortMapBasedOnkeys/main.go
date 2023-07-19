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

	sort.Ints(res)

	fmt.Println("After Sorting ", res)

	for _, k := range res {
		fmt.Println(k, m[k])
	}
}
