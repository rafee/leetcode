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

import "sync"

// @lc code=start
func solve(board [][]byte) {
	X := len(board)
	if X == 0 {
		return
	}
	Y := len(board[0])
	for i := 0; i < X; i++ {
		if board[i][0] == 'O' {
			visit(board, i, 0)
		}
		if board[i][Y-1] == 'O' {
			visit(board, i, Y-1)
		}
	}
	for j := 1; j < Y-1; j++ {
		if board[0][j] == 'O' {
			visit(board, 0, j)
		}
		if board[X-1][j] == 'O' {
			visit(board, X-1, j)
		}
	}

	wg := sync.WaitGroup{}
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			wg.Add(1)
			restoreBoard(board, i, j, &wg)
		}
	}
	wg.Wait()
}

func restoreBoard(board [][]byte, i int, j int, wg *sync.WaitGroup) {
	if board[i][j] == '*' {
		board[i][j] = 'O'
	} else if board[i][j] == 'O' {
		board[i][j] = 'X'
	}
	wg.Done()
}

func visit(board [][]byte, x, y int) {
	board[x][y] = '*'
	if x > 1 && board[x-1][y] == 'O' {
		visit(board, x-1, y)
	}
	if x < len(board)-2 && board[x+1][y] == 'O' {
		visit(board, x+1, y)
	}
	if y > 1 && board[x][y-1] == 'O' {
		visit(board, x, y-1)
	}
	if y < len(board[0])-2 && board[x][y+1] == 'O' {
		visit(board, x, y+1)
	}
}

// @lc code=end
