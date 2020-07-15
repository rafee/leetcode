/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

package leetcode

import "strings"

// @lc code=start
func reverseWords(s string) string {
	var sb strings.Builder
	segs := strings.Split(s, " ")
	for j := len(segs) - 1; j >= 0; j-- {
		if segs[j] == "" {
			continue
		}
		sb.WriteString(segs[j])
		sb.WriteString(" ")
	}
	if sb.Len() != 0 {
		return sb.String()[:sb.Len()-1]
	}
	return sb.String()
}

// @lc code=end
