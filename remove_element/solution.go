package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	k := 3
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// k := 2
	res := removeElement(nums, k)
	fmt.Println("res: ", res)
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
