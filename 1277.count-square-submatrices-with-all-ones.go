/*
 * @lc app=leetcode id=1277 lang=golang
 *
 * [1277] Count Square Submatrices with All Ones
 */

package leetcode

// @lc code=start
func countSquares(matrix [][]int) int {
	X, Y := len(matrix), len(matrix[0])
	count := 0

	for i := 0; i < X; i++ {
		count += matrix[i][0]
	}
	for j := 1; j < Y; j++ {
		count += matrix[0][j]
	}
	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			if matrix[i][j] != 1 {
				continue
			}
			matrix[i][j] = min3(matrix[i-1][j], matrix[i][j-1],
				matrix[i-1][j-1]) + 1
			count += matrix[i][j]
		}
	}
	return count
}

func min3(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
	} else if y < z {
		return y
	}
	return z
}

// @lc code=end
