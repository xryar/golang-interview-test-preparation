package main

import (
	"fmt"
	"strings"
)

func WordFrequency(s string) map[string]int {
	words := strings.Fields(s)

	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	return freq
}

func main() {
	sentence := "aku suka golang dan aku suka coding"
	result := WordFrequency(sentence)
	fmt.Println(result)
}
