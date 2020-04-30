/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (25.26%)
 * Likes:    1199
 * Dislikes: 548
 * Total Accepted:    188.3K
 * Total Submissions: 739K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 *
 */

package leetcode

// @lc code=start
func solve(board [][]byte) {
	nX := len(board)
	if nX == 0 {
		return
	}
	nY := len(board[0])
	for i := 0; i < nX; i++ {
		if board[i][0] == 'O' {
			visit(board, i, 0, nX, nY)
		}
		if board[i][nY-1] == 'O' {
			visit(board, i, nY-1, nX, nY)
		}
	}
	for j := 1; j < nY-1; j++ {
		if board[0][j] == 'O' {
			visit(board, 0, j, nX, nY)
		}
		if board[nX-1][j] == 'O' {
			visit(board, nX-1, j, nX, nY)
		}
	}
	for i := 0; i < nX; i++ {
		for j := 0; j < nY; j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func visit(board [][]byte, x, y, X, Y int) {
	if x < 0 || x > X-1 || y < 0 || y > Y-1 || board[x][y] != 'O' {
		return
	}

	board[x][y] = '*'
	visit(board, x-1, y, X, Y)
	visit(board, x+1, y, X, Y)
	visit(board, x, y-1, X, Y)
	visit(board, x, y+1, X, Y)
}

// @lc code=end
