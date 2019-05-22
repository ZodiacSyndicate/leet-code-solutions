/*
 * @lc app=leetcode.cn id=106 lang=javascript
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (60.75%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    7.9K
 * Total Submissions: 13K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
var buildTree = function(inorder, postorder) {
  function build(il, ir, pl, pr) {
    if (il > ir || pl > pr) return null
    const node = new TreeNode(postorder[pr])
    let index = 0
    for (let i = il; i <= ir; i++) {
      if (postorder[pr] === inorder[i]) {
        index = i
        break
      }
    }
    const rightCount = ir - index
    node.left = build(il, index - 1, pl, pr - rightCount - 1)
    node.right = build(index + 1, ir, pr - rightCount, pr - 1)
    return node
  }
  return build(0, inorder.length - 1, 0, postorder.length - 1)
}
