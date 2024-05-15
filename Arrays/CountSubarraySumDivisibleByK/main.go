package main

// Generate all the subarray and check how many
// subarrays having sum divisible by k then return count
// TC - O(N3)
func countSubArraySumDivisibleByKBrute(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]

			}
			if sum%k == 0 {
				count++
			}
		}
	}
	return count
}

// If we carefully observe, we can notice that to get the sum of the current subarray
// we just need to add the current element(i.e. arr[j]) to the sum of the previous subarray.
// we can remove the third loop and while moving j pointer, we can calculate the sum
// TC - O(N2)
func countSubArraySumDivisibleByKBetter(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum%k == 0 {
				count++
			}
		}
	}
	return count
}

// Take a map that will store the remainder and frequency
// Check in the map if there already exist the remainder, if yes
// then we will fetch its frequency and add into the count else
// simply put the element in map increasing its occurrence by 1
func countSubArraySumDivisibleByKOptimal(arr []int, k int) int {
	count := 0
	prefixSum := 0
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		prefixSum += arr[i]
		rem := prefixSum % k
		if rem < 0 {
			rem += k
		}
		count += m[rem]
		m[rem] += 1
	}
	return count

}
