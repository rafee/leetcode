/*
 * @lc app=leetcode id=107 lang=typescript
 *
 * [107] Binary Tree Level Order Traversal II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function levelOrderBottom(root: TreeNode | null): number[][] {
    let trees: (TreeNode | null)[] = [root]
    return helpLevelOrderBottomRecursive(trees).slice(1)
};

function helpLevelOrderBottomRecursive(trees: (TreeNode | null)[]): number[][] {
    if (trees.length == 0) {
        return []
    }
    let nextTrees: (TreeNode | null)[] = []
    let nums: number[] = []
    trees.forEach(node => {
        if (node !== null) {
            nums.push(node.val)
            nextTrees.push(node.left, node.right)
        }
    })
    let result = helpLevelOrderBottomRecursive(nextTrees)
    result.push(nums)
    return result
}
// @lc code=end

