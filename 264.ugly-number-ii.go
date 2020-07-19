/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

package golang

// @lc code=start
func nthUglyNumber(n int) int {
	uglies := make([]int, n)
	uglies[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num2 := uglies[i2] * 2
		num3 := uglies[i3] * 3
		num5 := uglies[i5] * 5
		minVal := min(num2, num3, num5)
		uglies[i] = minVal
		if minVal == num2 {
			i2++
		}
		if minVal == num3 {
			i3++
		}
		if minVal == num5 {
			i5++
		}
	}
	return uglies[n-1]
}

// @lc code=end
