package main

import (
	"fmt"
	"unicode"
)

func cleanString(s string) string {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	return string(cleaned)
}

func isPalindrome(s string) bool {
	s = cleanString(s)
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
	fmt.Println(isPalindrome("kasur ini rusak"))
}
