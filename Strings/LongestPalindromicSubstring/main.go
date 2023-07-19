// The program uses two nested loops to iterate over all possible substrings
// of the input string s, and checks if each substring is a palindrome.
// It keeps track of the longest palindrome found so far in the longest variable.
// The program first checks for palindromes with odd length, and then for palindromes
// with even length,

package main

import "fmt"

func LongestPalindromicSubstring(s string) string {
	n := len(s)
	longest := ""

	for i := 0; i < n; i++ {
		l, r := i, i

		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > len(longest) {
				longest = s[l : r+1]
			}
			l--
			r++
		}

		l, r = i, i+1

		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > len(longest) {
				longest = s[l : r+1]
			}
			l--
			r++
		}

	}

	return longest

}

func main() {
	str := "babad"
	fmt.Println(LongestPalindromicSubstring(str))
}
