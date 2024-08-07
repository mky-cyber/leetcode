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
	left := 0
	for right := 1; right < len(nums); right++ {
		// check for uniue element
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return left + 1
}
