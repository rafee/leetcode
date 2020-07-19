/*
 * @lc app=leetcode id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 */

package golang

import "math"

// @lc code=start
func maxSubarraySumCircular(A []int) int {
	maxSum, maxLSum, minSum, minLSum := math.MinInt32, math.MinInt32,
		0, 0
	totalSum := 0
	for _, num := range A {
		totalSum += num
		maxLSum = max(maxLSum+num, num)
		maxSum = max(maxSum, maxLSum)
		minLSum = min(minLSum+num, num)
		minSum = min(minSum, minLSum)
	}
	if totalSum-minSum == 0 {
		return maxSum
	}
	return max(maxSum, totalSum-minSum)
}

// @lc code=end
