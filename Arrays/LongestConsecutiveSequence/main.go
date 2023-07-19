package main

import "fmt"

func LongestConsecutiveSequence(arr []int) int {
	m := make(map[int]bool)
	longest := 1

	for _, nums := range arr {
		m[nums] = true
	}

	for k := range m {
		if !m[k-1] {
			count := 1

			for m[k+1] {
				k++
				count++
			}
			if count > longest {
				longest = count
			}
		}

	}
	return longest
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(LongestConsecutiveSequence(arr))
}
