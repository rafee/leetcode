/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

package golang

// @lc code=start
func setZeroes(matrix [][]int) {
	X, Y := len(matrix), len(matrix[0])
	// row, col := checkZero(X, Y, matrix)
	// setZero(X, Y, col, matrix, row)

	col0 := false
	for i := 0; i < X; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < Y; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := X - 1; i >= 0; i-- {
		for j := Y - 1; j >= 1; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

func setZero(X int, Y int, col []bool, matrix [][]int, row []bool) {
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func checkZero(X int, Y int, matrix [][]int) ([]bool, []bool) {
	row := make([]bool, X)
	col := make([]bool, Y)
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	return row, col
}

// @lc code=end
