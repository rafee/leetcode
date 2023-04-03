/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

package golang

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	// m := float64(coordinates[1][1]-coordinates[0][1]) /
	// 	float64(coordinates[1][0]-coordinates[0][0])
	x1, y1, x2, y2 := coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]
	for _, pos := range coordinates[2:] {
		if (pos[1]-y2)*(pos[0]-x1) != (pos[1]-y1)*(pos[0]-x2) {
			return false
		}
	}
	return true
}

// @lc code=end
