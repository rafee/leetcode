/*
 * @lc app=leetcode id=155 lang=typescript
 *
 * [155] Min Stack
 */

// @lc code=start
class MinStack {
    nums: number[] = []
    mins: number[] = []
    constructor() {}

    push(x: number): void {
        if (this.mins.length === 0 || x <= this.getMin()) {
            this.mins.push(x)
        }
        this.nums.push(x)
    }

    pop(): void {
        if (this.getMin() === this.top()) {
            this.mins.pop()
        }
        this.nums.pop()
    }

    top(): number {
        return this.nums[this.nums.length - 1]
    }

    getMin(): number {
        return this.mins[this.mins.length - 1]
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
// @lc code=end
