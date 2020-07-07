/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

package leetcode

// @lc code=start
func islandPerimeter(grid [][]int) int {
	s := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				conn := 4
				if j != 0 && grid[i][j-1] == 1 {
					conn--
				}
				if i != 0 && grid[i-1][j] == 1 {
					conn--
				}
				if j != len(grid[0])-1 && grid[i][j+1] == 1 {
					conn--
				}
				if i != len(grid)-1 && grid[i+1][j] == 1 {
					conn--
				}
				s += conn
			}
		}
	}
	return s
}

// @lc code=end
