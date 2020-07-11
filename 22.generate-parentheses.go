/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (58.56%)
 * Likes:    4395
 * Dislikes: 239
 * Total Accepted:    493.7K
 * Total Submissions: 816.3K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */

package leetcode

// @lc code=start
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	str := make([]byte, n*2)
	generate(&result, str, 0, 0, n)
	return result
}

func generate(result *[]string, str []byte, open, close, max int) {
	next := open + close
	if next == max*2 {
		*result = append(*result, string(str))
		return
	}
	if open < max {
		str[next] = '('
		generate(result, str, open+1, close, max)
	}
	if close < open {
		str[next] = ')'
		generate(result, str, open, close+1, max)
	}
}

// @lc code=end
