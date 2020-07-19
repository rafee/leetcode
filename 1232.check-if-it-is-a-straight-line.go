/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

package golang

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	m := float64(coordinates[1][1]-coordinates[0][1]) /
		float64(coordinates[1][0]-coordinates[0][0])
	x, y := coordinates[0][0], coordinates[0][1]
	for _, pos := range coordinates[2:] {
		if float64(pos[1]-y)/float64(pos[0]-x) != m {
			return false
		}
	}
	return true
}

// @lc code=end
