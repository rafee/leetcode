/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 *
 * https://leetcode.com/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (84.17%)
 * Likes:    1898
 * Dislikes: 336
 * Total Accepted:    413.9K
 * Total Submissions: 487.3K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 * You're given strings J representing the types of stones that are jewels, and
 * S representing the stones you have.  Each character in S is a type of stone
 * you have.  You want to know how many of the stones you have are also
 * jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are
 * letters. Letters are case sensitive, so "a" is considered a different type
 * of stone from "A".
 *
 * Example 1:
 *
 *
 * Input: J = "aA", S = "aAAbbbb"
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: J = "z", S = "ZZ"
 * Output: 0
 *
 *
 * Note:
 *
 *
 * S and J will consist of letters and have length at most 50.
 * The characters in J are distinct.
 *
 *
 */

package leetcode

// @lc code=start
func numJewelsInStones(J string, S string) int {
	set := make(map[rune]bool, 0)
	for _, jewel := range J {
		set[jewel] = true
	}
	count := 0
	for _, stone := range S {
		if set[stone] {
			count++
		}
	}
	return count
}

// @lc code=end
