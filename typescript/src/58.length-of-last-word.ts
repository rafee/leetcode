/*
 * @lc app=leetcode id=58 lang=typescript
 *
 * [58] Length of Last Word
 */

// @lc code=start
function lengthOfLastWord(s: string): number {
  let last = s.length - 1
  while (last >= 0 && s.charAt(last) === ' ') {
    last -= 1
  }
  let len = 0
  while (last >= 0 && s.charAt(last) !== ' ') {
    len += 1
    last -= 1
  }
  return len
}
// @lc code=end
