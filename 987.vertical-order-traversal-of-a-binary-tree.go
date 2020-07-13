/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
 */

package leetcode

// @lc code=start

func verticalTraversal(root *TreeNode) [][]int {
	left := make([][]int, 0)
	right := make([][]int, 0)
	if root != nil {
		helpVertical(root, &left, &right, 0)
	}
	result := make([][]int, 0)
	for i := len(left) - 1; i >= 0; i-- {
		result = append(result, left[i])
	}
	return append(result, right...)
}

func helpVertical(root *TreeNode, left, right *[][]int, pos int) {
	if pos >= 0 {
		if len(*right) == pos {
			*right = append(*right, []int{root.Val})
		} else {
			(*right)[pos] = append((*right)[pos], root.Val)
		}
	} else {
		index := -pos - 1
		if len(*left) == index {
			*left = append(*left, []int{root.Val})
		} else {
			(*left)[index] = append((*left)[index], root.Val)
		}
	}

	if root.Left != nil {
		helpVertical(root.Left, left, right, pos-1)
	}
	if root.Right != nil {
		helpVertical(root.Right, left, right, pos+1)
	}
}

// @lc code=end
