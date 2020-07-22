/*
 * @lc app=leetcode id=103 lang=typescript
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

function zigzagLevelOrder(root: TreeNode | null): number[][] {
  let result: number[][] = []
  helpZigzagLevelOrder({
    root: root,
    depth: 0,
    result: result,
  })
  return result
}
function helpZigzagLevelOrder({
  root,
  depth,
  result,
}: {
  root: TreeNode | null
  depth: number
  result: number[][]
}): void {
  if (root === null) {
    return
  }
  if (result.length === depth) {
    result.push([])
  }
  if (depth % 2 == 0) {
    result[depth].push(root.val)
  } else {
    result[depth].unshift(root.val)
  }
  helpZigzagLevelOrder({
    root: root.left,
    depth: depth + 1,
    result: result,
  })
  helpZigzagLevelOrder({
    root: root.right,
    depth: depth + 1,
    result: result,
  })
}
// @lc code=end
