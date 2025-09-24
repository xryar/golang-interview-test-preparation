package main

import (
	"fmt"
	"strconv"
)

func isPalindromeNumber(n int) bool {
	s := strconv.Itoa(n)
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromeNumberMath(n int) bool {
	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n = n / 10
	}

	return original == reversed
}

func main() {
	fmt.Println(isPalindromeNumber(121))
	fmt.Println(isPalindromeNumberMath(121))
}
