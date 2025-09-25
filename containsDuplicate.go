package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}

func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func findDuplicates(nums []int) []int {
	freq := make(map[int]int)
	result := []int{}

	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(findDuplicates([]int{1, 2, 3, 1, 2, 4}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
	fmt.Println(containsDuplicateSort([]int{1, 2, 3, 2}))
}
