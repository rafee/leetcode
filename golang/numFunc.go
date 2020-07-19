package golang

import "math"

func max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func popLast(nums []int) (int, []int) {
	num, nums := nums[len(nums)-1], nums[:len(nums)-1]
	return num, nums
}
