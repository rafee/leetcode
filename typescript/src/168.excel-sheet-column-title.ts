/*
 * @lc app=leetcode id=168 lang=typescript
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
function convertToTitle(n: number): string {
  let result = ''
  let num = n
  while (num > 0) {
    let digit = num % 26
    num = Math.floor(num / 26)
    if (digit === 0) {
      num -= 1
      digit = 26
    }
    const char = String.fromCharCode(64 + digit)
    result = char + result
  }
  return result
}
// @lc code=end
