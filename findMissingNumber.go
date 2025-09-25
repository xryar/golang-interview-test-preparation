package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums) + 1
	total := n * (n + 1) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return total - sum
}

func missingNumberXOR(nums []int) int {
	n := len(nums) + 1
	xorAll := 0

	for i := 1; i <= n; i++ {
		xorAll ^= i
	}
	for _, num := range nums {
		xorAll ^= num
	}

	return xorAll
}

func findMissing(nums []int, n int) []int {
	present := make([]bool, n+1)

	for _, num := range nums {
		if num <= n {
			present[num] = true
		}
	}

	missing := []int{}
	for i := 1; i <= n; i++ {
		if !present[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

func main() {
	nums := []int{1, 2, 4, 5, 8}
	fmt.Println(findMissing(nums, 8))

	nums = []int{2, 3, 7, 1, 8, 12, 20}
	fmt.Println(findMissing(nums, 12))

	nums = []int{4, 6, 7, 9}
	fmt.Println(findMissing(nums, 10))
}
