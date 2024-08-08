package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	// nums := []int{0,0,1,1,1,1,2,3,3}
	res := removeDuplicates(nums)
	fmt.Println("res: ", res)
}

func removeDuplicates(nums []int) int {
	left := 0
	i := 1
	for right := 1; right < len(nums); right++ {
		// check if we have the same element twice
		if nums[right] == nums[left] && i < 2 {
			// fmt.Println("found duplicate")
			// fmt.Println("left: ", left, " right: ", right)
			// fmt.Println("nums[left]: ", nums[left], " nums[right]: ", nums[right])
			i++
			left++
		}
		// if not, then we can copy the right number to left + 1
		if nums[right] != nums[left] {
			// fmt.Println("found unique number")
			// fmt.Println("left: ", left, " right: ", right)
			i = 1
			nums[left+1] = nums[right]
			// fmt.Println("nums[left]: ", nums[left], " nums[right]: ", nums[right])
			left++
		}
	}
	fmt.Println(nums)
	return left + 1
}
