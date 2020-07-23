/*
 * @lc app=leetcode id=396 lang=typescript
 *
 * [396] Rotate Function
 */

// @lc code=start
function maxRotateFunction(A: number[]): number {
  if (A.length === 0) {
    return 0
  }
  const sum = A.reduce((total, current) => total + current, 0)
  let curSum = 0
  for (let i = 1; i < A.length; i++) {
    curSum += A[i] * i
  }
  let maxSum = Math.max(Number.MIN_SAFE_INTEGER, curSum)
  for (let i = A.length - 1; i > 0; i--) {
    curSum += sum
    curSum -= A[i] * A.length
    maxSum = Math.max(curSum, maxSum)
  }
  return maxSum
}
// @lc code=end
