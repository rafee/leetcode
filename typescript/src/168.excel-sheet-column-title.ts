/*
 * @lc app=leetcode id=168 lang=typescript
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
function convertToTitle(n: number): string {
    let result = ''
    while (n > 0) {
        let digit = n % 26
        n = Math.floor(n / 26)
        if (digit === 0) {
            n--
            digit = 26
        }
        let char = String.fromCharCode(64 + digit)
        result = char + result
    }
    return result
}
// @lc code=end
