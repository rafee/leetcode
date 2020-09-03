/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

package golang

// @lc code=start

// BSTIterator iterates over the binary search Trees
type BSTIterator struct {
	stack []*TreeNode
}

func bstConstructor(root *TreeNode) BSTIterator {
	iter := new(BSTIterator)
	iter.stack = make([]*TreeNode, 0)
	treeIter := root
	for treeIter != nil {
		iter.stack = append(iter.stack, treeIter)
		treeIter = treeIter.Left
	}
	return *iter
}

/** @return the next smallest number */

// Next returns the next smallest number
func (this *BSTIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	push := top.Right
	for push != nil {
		this.stack = append(this.stack, push)
		push = push.Left
	}
	return top.Val
}

/** @return whether we have a next smallest number */

//HasNext return whether the next iterator is present or not
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
