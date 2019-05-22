/*
 * @lc app=leetcode.cn id=95 lang=javascript
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (54.73%)
 * Likes:    131
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 10K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
 * @param {number} n
 * @return {TreeNode[]}
 */
var generateTrees = function(n) {
  if (n == 0) return []
  return generate(1, n)
}

function generate(start, end) {
  const res = []
  if (start > end) {
    res.push(null)
    return res
  }
  for (let i = start; i <= end; i++) {
    const subLeft = generate(start, i - 1)
    const subRight = generate(i + 1, end)
    for (const left of subLeft) {
      for (const right of subRight) {
        const node = new TreeNode(i)
        node.left = left
        node.right = right
        res.push(node)
      }
    }
  }
  return res
}
