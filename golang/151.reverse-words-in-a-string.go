/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

package golang

import "strings"

// @lc code=start
func reverseWords(s string) string {
	segs := strings.Fields(s)
	for i := 0; i < len(segs)/2; i++ {
		segs[i], segs[len(segs)-i-1] = segs[len(segs)-i-1], segs[i]
	}
	return strings.Join(segs, " ")
}

// @lc code=end
