package main

import "fmt"

func countDuplicateInArray(arr []int) int {

	freqMap := make(map[int]int)
	count := 0

	for _, value := range arr {
		freqMap[value] = freqMap[value] + 1
	}

	for _, v := range freqMap {
		if v > 1 {
			count = count + 1
		}
	}
	return count
}

func main() {

	arr := []int{12, 21, 12, 32, 44, 32, 88, 28, 44}
	fmt.Println(countDuplicateInArray(arr))

}
