package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 1, 2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println("res: ", res)
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
