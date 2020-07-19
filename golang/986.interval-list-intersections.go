/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

package golang

// @lc code=start
func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		start, end := getInterval(A[i], B[j])
		if start <= end {
			result = append(result, []int{start, end})
		}
		if A[i][1] == end {
			i++
		}
		if B[j][1] == end {
			j++
		}
	}
	return result
}

func getInterval(A []int, B []int) (int, int) {
	start := max(A[0], B[0])
	end := min(A[1], B[1])
	return start, end
}

// @lc code=end
