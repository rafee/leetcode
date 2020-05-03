/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 *
 * https://leetcode.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (51.15%)
 * Likes:    492
 * Dislikes: 172
 * Total Accepted:    153.5K
 * Total Submissions: 293K
 * Testcase Example:  '"a"\n"b"'
 *
 *
 * Given an arbitrary ransom note string and another string containing letters
 * from all the magazines, write a function that will return true if the
 * ransom
 * note can be constructed from the magazines ; otherwise, it will return
 * false.
 *
 *
 * Each letter in the magazine string can only be used once in your ransom
 * note.
 *
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 *
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 *
 */

package leetcode

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	set := make(map[rune]int, 0)
	for _, char := range magazine {
		set[char]++
	}

	for _, char := range ransomNote {
		val, ok := set[char]
		if ok && val > 0 {
			set[char]--
		} else {
			return false
		}
	}

	return true
}

// @lc code=end
