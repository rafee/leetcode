/*
 * @lc app=leetcode id=1779 lang=golang
 *
 * [1779] Find Nearest Point That Has the Same X or Y Coordinate
 */

package golang

import (
	"math"
)

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	distance, index := math.MaxInt64, -1
	for i, point := range points {
		if point[0] == x {
			dist := diff(point[1], y)
			if dist < distance {
				distance = dist
				index = i
			}
		} else if point[1] == y {
			x1 := diff(point[0], x)
			if x1 < distance {
				distance = x1
				index = i
			}
		}
	}
	return index
}

// @lc code=end
