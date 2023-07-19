package main

import "fmt"

func findFrequencyOfAnArrayElement(arr []int, num int) int {

	count := 0

	for _, value := range arr {
		if value == num {
			count = count + 1
		}
	}
	return count
}

func main() {

	arr := []int{12, 43, 22, 12, 2, 4, 255, 12, 43}
	num := 12

	fmt.Println(findFrequencyOfAnArrayElement(arr, num))

}
