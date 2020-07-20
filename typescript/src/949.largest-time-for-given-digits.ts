/*
 * @lc app=leetcode id=949 lang=typescript
 *
 * [949] Largest Time for Given Digits
 */

// @lc code=start
function largestTimeFromDigits(A: number[]): string {
    let maxHour = -1, maxMinute = -1
    for (let i = 0; i < A.length; i++) {
        for (let j = 0; j < A.length; j++) {
            if (i === j) {
                continue
            }
            let localMaxMinute = -1
            let localMaxHour = -1
            for (let k = 0; k < A.length; k++) {
                if (i === k || j === k) {
                    continue
                }
                for (let l = 0; l < A.length; l++) {
                    if (i === l || j === l || k === l) {
                        continue
                    }
                    const hour = A[i] * 10 + A[j]
                    const minute = A[k] * 10 + A[l]
                    if (hour >= localMaxHour && hour < 24) {
                        if (minute >= localMaxMinute && minute < 60) {
                            localMaxHour = hour
                            localMaxMinute = minute
                        }
                    }
                }
            }
            if (localMaxHour > maxHour) {
                maxHour = localMaxHour
                maxMinute = localMaxMinute
            }
        }
    }
    if (maxHour == -1 || maxMinute == -1) {
        return ""
    }
    let hourString = (`00${maxHour}`).slice(-2)
    let minuteString = (`00${maxMinute}`).slice(-2)
    return `${hourString}:${minuteString}`
};
// @lc code=end

