/*
 * @lc app=leetcode id=27 lang=typescript
 *
 * [27] Remove Element
 */

// @lc code=start
function removeElement(nums: number[], val: number): number {
    let p = 0;
    for (const v of nums) {
        if (v !== val) {
            nums[p++] = v;
        }
    }
    return p;
};

// @lc code=end

