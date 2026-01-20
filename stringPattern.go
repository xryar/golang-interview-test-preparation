package main

import "fmt"

func printPattern(s1, s2 string) {
	n := len(s1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				fmt.Print(string(s1[i]))
			} else if j == n-1-i {
				fmt.Print(string(s2[i]))
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

func main() {
	printPattern("coder", "byte")
}
