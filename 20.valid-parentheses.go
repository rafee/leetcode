/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.72%)
 * Likes:    3803
 * Dislikes: 190
 * Total Accepted:    792.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

package golang

// @lc code=start

func isValid(s string) bool {
	stack := make([]byte, 0)
	lookupMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack = append(stack, lookupMap[s[i]])

		case ')':
			fallthrough
		case '}':
			fallthrough
		case ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != s[i] {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end
