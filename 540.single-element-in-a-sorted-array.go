/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

package leetcode

// @lc code=start
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] != nums[mid+1] {
			right = mid
		} else {
			left = mid + 2
		}
	}
	return nums[left]
}

// @lc code=end
