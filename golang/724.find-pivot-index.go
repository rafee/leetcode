/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

package golang

// @lc code=start
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for i, runningSum := 0, 0; i < len(nums); i++ {
		if runningSum == sum-nums[i]-runningSum {
			return i
		}
		runningSum += nums[i]
	}
	return -1
}

// @lc code=end
