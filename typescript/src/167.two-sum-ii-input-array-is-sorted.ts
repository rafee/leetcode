/*
 * @lc app=leetcode id=167 lang=typescript
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
function twoSumSorted(numbers: number[], target: number): number[] {
    const numMap: Map<number, number> = new Map()
    for (let i = 0; i < numbers.length; i++) {
        if (numMap.has(numbers[i])) {
            return [numMap.get(numbers[i])!, i + 1]
            // } else if (numbers[i] > target - numbers[i - 1]) {
            //     break
        } else {
            numMap.set(target - numbers[i], i + 1)
        }
    }
    return []
}
// @lc code=end
