/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (34.71%)
 * Likes:    2843
 * Dislikes: 127
 * Total Accepted:    447K
 * Total Submissions: 1.3M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 */

package golang

// @lc code=start
func searchRange(nums []int, target int) []int {
	index := -1
	for start, end := 0, len(nums)-1; start <= end; {
		mid := start + (end-start)/2
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			index = mid
			break
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	rightIndex := 0
	for start, end := index, len(nums)-1; start <= end; {
		mid := start + (end-start)/2
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
			rightIndex = mid
			break
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return []int{index, rightIndex}

	// for start, end := 0, len(nums)-1;  < count; i++ {

	// }
}

// @lc code=end
