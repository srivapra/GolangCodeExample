// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

// Using Extra space
func merge1(nums1 []int, m int, nums2 []int, n int) {

	arr3 := make([]int, len(nums1)+len(nums2))

	for i := 0; i < m; i++ {
		arr3[i] = nums1[i]
	}

	for i := 0; i < n; i++ {
		arr3[i+m] = nums2[i]
	}
	sort.Ints(arr3[:])
	fmt.Print(arr3)

}

// Without Extra space
func merge2(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1, 0, 0, 0)

	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}

	fmt.Print(nums1)

}

func main() {
	arr1 := []int{2, 4, 7, 7, 8, 12}
	m := len(arr1)
	arr2 := []int{1, 3, 6}
	n := len(arr2)

	merge1(arr1, m, arr2, n)
	merge2(arr1, m, arr2, n)
}
