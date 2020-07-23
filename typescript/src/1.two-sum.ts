/*
 * @lc app=leetcode id=1 lang=typescript
 *
 * [1] Two Sum
 */

// @lc code=start
function twoSum(nums: number[], target: number): number[] {
  const dict: Map<number, number> = new Map()

  for (let i = 0; i < nums.length; i++) {
    const preLoc = dict.has(nums[i])
    if (!preLoc) {
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      return [dict.get(nums[i])!, i]
    }
    dict.set(target - nums[i], i)
  }

  return []
}
// @lc code=end
