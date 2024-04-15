package main

/* Brute Force Approach
1. Put the first k element in temp array
2. Shift the element start from the k till end by one place before i.e : arr[i-k] = arr[i]
3. Put back the temp element array in original array
*/
import "fmt"

// Reverse the first k elements.
// Reverse the remaining n - k elements.
// Reverse the entire array.

func rotateLeft(nums []int, k int) {
	n := len(nums)
	k = k % n // Ensure k is within the array length
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
	reverse(nums, 0, n-1)
	fmt.Println(nums)

}

func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n             // Ensure k is within the array length
	reverse(nums, 0, n-1) // Reverse the entire array
	reverse(nums, 0, k-1) // Reverse the first k elements
	reverse(nums, k, n-1) // Reverse the remaining elements
	fmt.Println(nums)

}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 8
	rotateLeft(arr, k)
}
