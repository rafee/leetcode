/*
 * @lc app=leetcode id=1523 lang=golang
 *
 * [1523] Count Odd Numbers in an Interval Range
 */

package golang

// @lc code=start
func countOdds(low int, high int) int {
	count := 0
	if high%2 == 1 || low%2 == 1 {
		count++
	}
	return (high-low)/2 + count
}

// @lc code=end
