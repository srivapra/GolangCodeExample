package main

import "fmt"

func isPalindromeNumber(num int) bool {

	inputNumber := num

	reverseNumber := 0

	for num > 0 {

		reverseNumber = reverseNumber*10 + num%10
		num = num / 10
	}

	if reverseNumber == inputNumber {
		return true
	} else {
		return false
	}
}

func main() {
	num := 121
	fmt.Println("Input Number : ", num)
	fmt.Println(isPalindromeNumber(num))
}
