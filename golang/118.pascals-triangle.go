/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

package golang

// @lc code=start
func genPascal(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = row
	}
	return res
}

// @lc code=end
