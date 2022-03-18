package main

import (
	"fmt"
	"math"
)

func main() {
	nums := make([]int, 5)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = -1

	fmt.Printf("Max contiguous subarray sum: %d", kagane(nums))
}

// Kagane's algorithm is used to find the sum of the maximum continuous subarray
func kagane(nums []int) int {
	local_max := 0
	global_max := math.MinInt

	for i := 0; i < len(nums); i++ {
		local_max = max(nums[i], local_max+nums[i])
		if local_max > global_max {
			global_max = local_max
		}
	}

	return global_max

}

// Helper function to return the max of two integers, since the built-in
// function in the math package is for floats
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
