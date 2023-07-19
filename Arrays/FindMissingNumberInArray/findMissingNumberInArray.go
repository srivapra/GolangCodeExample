package main

import "fmt"

func findMissingNumberInArray(arr []int) int {

	n := len(arr)

	arraySum := 0

	sumOfNaturalNumber := ((n + 1) * (n + 2)) / 2

	for i := 0; i < n; i++ {
		arraySum = arraySum + arr[i]

	}

	return sumOfNaturalNumber - arraySum
}

func main() {

	arr := []int{1, 2, 4, 6, 3, 7, 8}
	fmt.Println(findMissingNumberInArray(arr))

}
