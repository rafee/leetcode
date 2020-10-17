/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

package golang

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	r, found := getRowLoc(matrix, target)
	if found {
		found = getColLoc(matrix[r], target)
	}
	return found
}

func getRowLoc(matrix [][]int, target int) (int, bool) {
	X := len(matrix)
	if X == 0 {
		return 0, false
	}

	for low, high := 0, X-1; low <= high; {
		mid := low + (high-low)/2
		if matrix[mid][0] > target {

		}
	}
}

// @lc code=end
