/*
 * @lc app=leetcode id=775 lang=golang
 *
 * [775] Global and Local Inversions
 */

package golang

// @lc code=start
func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i]-i > 1 || i-A[i] > 1 {
			return false
		}
	}
	return true
}

// @lc code=end
