package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum]++
	}
	return count
}

func main() {
	arr := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(arr, k))

}
