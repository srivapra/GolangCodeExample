package main

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func expandAroundCenter(str string, left int, right int) int {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return right - left - 1
}

// Optimial approach - O(n^2)
func longestPalindromeSubsOp(str string) string {
	var start, end int

	for i := 0; i < len(str); i++ {
		len1 := expandAroundCenter(str, i, i)
		len2 := expandAroundCenter(str, i, i+1)
		length := len1
		if len2 > len1 {
			length = len2
		}
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2

		}
	}
	return str[start : end+1]
}

// Naive Approach - generate all the substring and check the substring
// is palindrome or not and update the longest substring - O (n^3)
func longestPalindromeSubs(str string) string {
	longest := ""
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			substring := str[i:j]
			if isPalindrome(substring) && len(substring) > len(longest) {
				longest = substring
			}
		}
	}
	return longest
}

func main() {
	str := "Geeks"
	fmt.Println(longestPalindromeSubs(str))
}
