/*
 * @lc app=leetcode id=27 lang=typescript
 *
 * [27] Remove Element
 */

// @lc code=start
function removeElement(nums: number[], val: number): number {
    let end = nums.length
    for (let i = 0; i < end; ) {
        if (nums[i] === val) {
            nums[i] = nums[end - 1]
            end--
        } else {
            i++
        }
    }
    return end
}

// @lc code=end
