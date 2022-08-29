/*
 * @lc app=leetcode id=1491 lang=golang
 *
 * [1491] Average Salary Excluding the Minimum and Maximum Salary
 */

package golang

// @lc code=start
func average(salary []int) float64 {
	max, min, sum := 0, 10000000, 0
	for _, v := range salary {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}

// @lc code=end
