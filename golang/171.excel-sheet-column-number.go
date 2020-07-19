/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

package golang

// @lc code=start
func titleToNumber(s string) int {
	col := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i]-'A') + 1
		col *= 26
		col += c
	}
	return col
}

// @lc code=end
