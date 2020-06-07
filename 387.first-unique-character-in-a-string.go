/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (51.12%)
 * Likes:    1696
 * Dislikes: 112
 * Total Accepted:    455.4K
 * Total Submissions: 868.9K
 * Testcase Example:  '"leetcode"'
 *
 *
 * Given a string, find the first non-repeating character in it and return it's
 * index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 *
 *
 *
 *
 * Note: You may assume the string contain only lowercase letters.
 *
 */

package leetcode

// @lc code=start
func firstUniqChar(s string) int {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return i
		}
	}

	return -1
}

// @lc code=end
