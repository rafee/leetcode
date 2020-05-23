/*
 * @lc app=leetcode id=329 lang=golang
 *
 * [329] Longest Increasing Path in a Matrix
 */

package leetcode

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	X := len(matrix)
	if X == 0 {
		return 0
	}
	Y := len(matrix[0])
	seqMatrix := make([][]int, X)
	for i := 0; i < X; i++ {
		seqMatrix[i] = make([]int, Y)
	}

	var maxL int
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			checkLongSeq(matrix, seqMatrix, i, j, X, Y)
			if seqMatrix[i][j] > maxL {
				maxL = seqMatrix[i][j]
			}
		}
	}
	return maxL
}

func checkLongSeq(matrix [][]int, seqMatrix [][]int, i int, j int, X int, Y int) {
	if seqMatrix[i][j] != 0 {
		return
	}
	maxVal := 1
	for _, next := range nextValid(i, j, X, Y) {
		if matrix[next[0]][next[1]] > matrix[i][j] {
			checkLongSeq(matrix, seqMatrix, next[0], next[1], X,
				Y)
			nextVal := seqMatrix[next[0]][next[1]] + 1
			if nextVal > maxVal {
				maxVal = nextVal
			}
		}
	}
	seqMatrix[i][j] = maxVal
}

func nextValid(i, j, X, Y int) [][]int {
	result := make([][]int, 0)
	if i+1 < X {
		result = append(result, []int{i + 1, j})
	}
	if i-1 >= 0 {
		result = append(result, []int{i - 1, j})
	}
	if j+1 < Y {
		result = append(result, []int{i, j + 1})
	}
	if j-1 >= 0 {
		result = append(result, []int{i, j - 1})
	}
	return result
}

// @lc code=end
