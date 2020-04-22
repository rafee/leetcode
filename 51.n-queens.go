/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (42.74%)
 * Likes:    1547
 * Dislikes: 65
 * Total Accepted:    185.5K
 * Total Submissions: 417.9K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

package leetcode

// @lc code=start
func solveNQueens(n int) [][]string {
	board := makeByteBoard(n)
	return solveNQueensStr(board, n, 0, []string{})
}

func makeByteBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return board
}

func solveNQueensStr(board [][]byte, n, row int, solution []string) [][]string {
	result := make([][]string, 0)
	for col := 0; col < n; col++ {
		if isSafeStr(board, n, row, col) {
			sol := placeQueenStr(board, n, row, col)
			tmp := append(solution, sol)
			if row != (n - 1) {
				res := solveNQueensStr(board, n, row+1, tmp)
				result = append(result, res...)
			} else {
				result = append(result, tmp)
			}
			removeQueenStr(board, n, row, col)
		}
	}
	return result
}

func isSafeStr(board [][]byte, n, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func placeQueenStr(board [][]byte, n, row, col int) string {
	board[row][col] = 'Q'
	return string(board[row])
}

func removeQueenStr(board [][]byte, n, row, col int) {
	board[row][col] = '.'
}

// @lc code=end
