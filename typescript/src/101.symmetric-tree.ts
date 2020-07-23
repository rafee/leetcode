/*
 * @lc app=leetcode id=101 lang=typescript
 *
 * [101] Symmetric Tree
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

function isSymmetric(root: TreeNode | null): boolean {
  if (root === null) {
    return true;
  }
  return helpIsSymmetricRecursive({ left: root.left, right: root.right });
}

function helpIsSymmetricRecursive({
  left,
  right,
}: {
  left: TreeNode | null
  right: TreeNode | null
}): boolean {
  if (!left || !right) {
    return left == right;
  }
  if (left.val !== right.val) {
    return false;
  }
  return (
    helpIsSymmetricRecursive({ left: left.left, right: right.right })
        && helpIsSymmetricRecursive({ left: left.right, right: right.left })
  );
}
// @lc code=end
