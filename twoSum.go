package main

import "fmt"

func twoSums(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := seen[complement]; found {
			return []int{idx, i}
		}

		seen[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	result := twoSums(nums, target)
	fmt.Println(result)
}
