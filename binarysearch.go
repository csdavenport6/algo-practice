package main

import (
	"algopractice/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fp := "sorted_input.txt"
	f, err := os.Open(fp)
	helpers.Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	nums := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		helpers.Check(err)
		nums = append(nums, i)
	}

	fmt.Printf("%d\n", binarysearch(nums, 9))
}

func binarysearch(nums []int, target int) int {
	low, high := 0, len(nums)

	for low <= high {
		mid := low + (high-low)/2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
