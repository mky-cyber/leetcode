package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	res := merge(nums1, m, nums2, n)
	fmt.Println("Sorted: ", res)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 && n > 0 {
		copy(nums1, nums2)
	}

	if m > 0 && n > 0 {
		nums1 = nums1[:m]
		nums1 = append(nums1, nums2...)
		sort.Ints(nums1)
	}

	return nums1
}
