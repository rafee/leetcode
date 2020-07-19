/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

package golang

// @lc code=start
func hIndexII(citations []int) int {
	left, right := 0, len(citations)-1
	for left <= right {
		mid := left + (right-left)/2
		h := len(citations) - mid
		if citations[mid] == h {
			return h
		} else if citations[mid] < h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return len(citations) - left
}

// @lc code=end
