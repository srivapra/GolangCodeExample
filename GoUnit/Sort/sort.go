package main

import "fmt"

func SortArrayOf012(arr []int) []int {

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
	return arr
}

func main() {

	arr := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	fmt.Println(SortArrayOf012(arr))

}
