/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

package golang

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	x, y := len(mat), len(mat[0])
	if x*y != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		result[i/c][i%c] = mat[i/y][i%y]
	}
	return result
}

// @lc code=end
