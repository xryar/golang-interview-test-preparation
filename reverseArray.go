package main

import "fmt"

func reverseArray(nums []int) {
	i := 0
	j := len(nums) - 1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	arr := []int{1, 3, 2, 4, 5}
	reverseArray(arr)
	fmt.Println(arr)
}
