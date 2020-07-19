/*
 * @lc app=leetcode id=901 lang=golang
 *
 * [901] Online Stock Span
 */

package golang

// @lc code=start
type StockSpanner struct {
	stack [][2]int
}

func StockConstructor() StockSpanner {
	newStockSpanner := new(StockSpanner)
	newStockSpanner.stack = make([][2]int, 0)
	return *newStockSpanner
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	for {
		if len(this.stack) == 0 {
			break
		}
		curElem := this.stack[len(this.stack)-1]
		if curElem[0] > price {
			break
		}
		res += curElem[1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, res})
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end
