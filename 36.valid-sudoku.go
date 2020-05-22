/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

package leetcode

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	box := make([][]map[byte]bool, 3)
	for i := 0; i < 3; i++ {
		box[i] = make([]map[byte]bool, 3)
		for j := 0; j < 3; j++ {
			box[i][j] = make(map[byte]bool)
		}
	}

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			boxX, boxY := i/3, j/3
			if box[boxX][boxY][val] || row[i][val] || col[j][val] {
				return false
			}
			box[boxX][boxY][val] = true
			row[i][val] = true
			col[j][val] = true
		}
	}
	return true
}

// @lc code=end
