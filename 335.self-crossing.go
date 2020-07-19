/*
 * @lc app=leetcode id=335 lang=golang
 *
 * [335] Self Crossing
 */

package golang

// @lc code=start
func isSelfCrossing(x []int) bool {
	if len(x) < 4 {
		return false
	}
	for i := 0; i < len(x); i++ {
		chk := i % 4
		if chk < 3 && chk != 0 {
			x[i] = -x[i]
		}
	}
	var cur [4][2]int
	cur[0] = [2]int{0, 0}
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			cur[i+1][0] = cur[i][0]
			cur[i+1][1] = cur[i][1] + x[i]
		} else {
			cur[i+1][0] = cur[i][0] + x[i]
			cur[i+1][1] = cur[i][1]
		}
	}

	for i := 3; i < len(x); i++ {
		var nextPos [2]int
		if i%2 == 0 {
			nextPos[0] = cur[i][0]
			nextPos[1] = cur[i][1] + x[i]
		} else {
			nextPos[0] = cur[i][0] + x[i]
			nextPos[1] = cur[i][1]
		}
		if checkCross(nextPos, cur[3], cur[2], cur[1]) {
			return true
		}
		cur[2], cur[1], cur[0] = cur[3], cur[2], cur[1]
		cur[3] = nextPos
	}
	return false
}

func checkCross(pos1, pos2, pos3, pos4 [2]int) bool {
	if pos3[0] == pos4[0] {
		if sign(pos1[0]-pos3[0]) != sign(pos2[0]-pos3[0]) {
			return sign(pos3[1]-pos1[1]) != sign(pos4[1]-pos1[1])
		}
	} else {
		if sign(pos1[1]-pos3[1]) != sign(pos2[1]-pos3[1]) {
			return sign(pos3[0]-pos1[0]) != sign(pos4[0]-pos1[0])
		}
	}
	return false
}

func sign(n int) bool {
	return n > 0
}

// @lc code=end
