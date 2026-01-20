package main

import "fmt"

func findCommon(arr [][]int) string {
	result := []int{}

	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr[1]); j++ {
			if arr[0][i] == arr[1][j] {
				result = append(result, arr[0][i])
			}
		}
	}

	return fmt.Sprintf("%v asdwasdwa", result)
}

func main() {
	a := [][]int{{1, 2, 4, 6}, {1, 2, 3, 5, 7}}
	fmt.Println(findCommon(a))
}
