/*
 * @lc app=leetcode id=171 lang=typescript
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
function titleToNumber(s: string): number {
    let num = 0
    const a = "A".charCodeAt(0) - 1
    for (let i = 0; i < s.length; i++) {
        const element = s.charCodeAt(i);
        num *= 26
        num += (element - a)
    }
    return num
};
// @lc code=end

