package main

import (
	"fmt"
	"sort"
	"strings"
)

// use sort
func sortString(s string) string {
	s = strings.ToLower(s)
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return sortString(a) == sortString(b)
}

// map
func isAnagramMap(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range a {
		count[ch]++
	}

	for _, ch := range b {
		count[ch]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagramMap("silent", "listen"))
	fmt.Println(isAnagram("hello", "word"))
	fmt.Println(isAnagram("acumalaka", "malacukaa"))
}
