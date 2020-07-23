/*
 * @lc app=leetcode id=110 lang=typescript
 *
 * [110] Balanced Binary Tree
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

function isBalanced(root: TreeNode | null): boolean {
  const [, isBal] = helpIsBalanced(root);
  return isBal;
}

function helpIsBalanced(root: TreeNode | null): [number, boolean] {
  if (root == null) {
    return [0, true];
  }
  const [lDepth, lBal] = helpIsBalanced(root.left);
  const [rDepth, rBal] = helpIsBalanced(root.right);
  if (Math.abs(rDepth - lDepth) > 1) {
    return [0, false];
  }
  return [Math.max(lDepth, rDepth) + 1, lBal && rBal];
}
// @lc code=end
