package main

import "fmt"

func removeDuplicateFromSlice(s []string) []string {

	inResult := make(map[string]bool)

	var result []string

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// This function, however, needs to be reimplemented
// each time the slice is of a different type.
// So, if we had []int and []string slices that we wanted to remove duplicates from,
// so far, we needed two functions: uniqueString() and uniqueInt().
// But with the release of the Generics in Go 1.18, this is no longer necessary.
// We can write a single function that will work on any slice where the values
// satisfy the comparable constraint.
func removeDuplicateFromAnySlice[T comparable](s []T) []T {

	inResult := make(map[T]bool)

	var result []T

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {

	str := []string{"abc", "cde", "efg", "abc", "cde"}
	fmt.Println("Without Generics : ", removeDuplicateFromSlice(str))

	str1 := []string{"abc", "cde", "efg", "abc", "cde"}
	fmt.Println("By Generics : ", removeDuplicateFromAnySlice(str1))

	int := []int{1, 1, 2, 2, 3, 3, 4}
	fmt.Println("By Generics : ", removeDuplicateFromAnySlice(int))

}
