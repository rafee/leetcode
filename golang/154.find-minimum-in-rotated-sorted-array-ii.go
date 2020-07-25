/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

package golang

// @lc code=start
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	// l == r when we found our minumum value
	for l < r {
		mid := l + (r-l)/2

		if nums[l] > nums[mid] {
			// left side is unsorted, so our answer is inside left
			r = mid
		} else if nums[mid] > nums[r] {
			// right side is unsorted, so our answer is inside right.
			// +1 because mid will always larger than the minumum
			l = mid + 1
		} else if nums[mid] == nums[r] {
			// this is when this question get tricky.
			// think of the case like [2 2 2], [1, 3, 3], [3, 1, 3, 3]
			r = r - 1
		} else {
			// it's sorted
			r = l
		}
	}

	return nums[l]
}

// @lc code=end
