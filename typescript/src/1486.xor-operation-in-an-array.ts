/*
 * @lc app=leetcode id=1486 lang=typescript
 *
 * [1486] XOR Operation in an Array
 */

// @lc code=start
function xorOperation(n: number, start: number): number {
  let num = start
  for (let i = 1; i < n; i++) {
    // start += 2;
    // eslint-disable-next-line no-bitwise
    num ^= start + 2 * i
  }
  return num
}
// @lc code=end
