/*
 * @lc app=leetcode id=1672 lang=golang
 *
 * [1672] Richest Customer Wealth
 */

package golang

// @lc code=start
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, row := range accounts {
		s := 0
		for _, v := range row {
			s += v
		}
		if s > max {
			max = s
		}
	}
	return max
}

// @lc code=end
