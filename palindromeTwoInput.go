package main

import (
	"fmt"
	"strconv"
)

func isPalindromeTwoInput(input interface{}) bool {
	var s string

	switch v := input.(type) {
	case string:
		s = v
	case int:
		s = strconv.Itoa(v)
	default:
		return false
	}

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

func main() {
	fmt.Println(isPalindromeTwoInput(121))
	fmt.Println(isPalindromeTwoInput("aba"))
}
