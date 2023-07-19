package main

import "fmt"

func sortArrayOf012(arr []int) {

	arrLength := len(arr)

	low := 0
	mid := 0
	high := arrLength - 1

	for mid <= high {
		switch arr[mid] {
		case 0:
			{
				arr[low], arr[mid] = arr[mid], arr[low]
				low++
				mid++
				break
			}
		case 1:
			{
				mid++
				break
			}
		case 2:
			{
				arr[mid], arr[high] = arr[high], arr[mid]
				high--
				break

			}
		}

	}
}

func printSortedArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func main() {

	arr := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	sortArrayOf012(arr)
	printSortedArray(arr)

}
