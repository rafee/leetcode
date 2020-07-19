/*
 * @lc app=leetcode id=58 lang=typescript
 *
 * [58] Length of Last Word
 */

// @lc code=start
function lengthOfLastWord(s: string): number {
    let last = s.length - 1
    while (last >= 0 && s.charAt(last) === " ") {
        last--
    }
    let len = 0
    while (last >= 0 && s.charAt(last) !== " ") {
        len++
        last--
    }
    return len
};
// @lc code=end

