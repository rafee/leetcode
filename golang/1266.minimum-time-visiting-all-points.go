/*
 * @lc app=leetcode id=1266 lang=golang
 *
 * [1266] Minimum Time Visiting All Points
 */

package golang

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	dist := 0
	for i := 1; i < len(points); i++ {
		x_diff := diff(points[i][0], points[i-1][0])
		y_diff := diff(points[i][1], points[i-1][1])
		dist += max(x_diff, y_diff)
	}
	return dist
}

func diff(val1, val2 int) int {
	if val1 > val2 {
		return val1 - val2
	}
	return val2 - val1
}

// @lc code=end
