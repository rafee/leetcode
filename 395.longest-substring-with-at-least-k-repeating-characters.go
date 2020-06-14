/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

package leetcode

// @lc code=start
func longestSubstring(s string, k int) int {
	cMap := make(map[byte]int)
	mx := 0
	for i, start, r := 0, 0, 0; i < len(s); i++ {
		ch := s[i]
		if cMap[ch] < start {
			r++
			cMap[ch] = i
		}
		l := i - start
		if l > mx {
			mx = l
		}
		if r == k {
			start = cMap[s[i-1]] + 1
		}
	}
	return mx
}

// @lc code=end
