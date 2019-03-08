/*
 * @lc app=leetcode.cn id=427 lang=javascript
 *
 * [427] Construct Quad Tree
 *
 * https://leetcode-cn.com/problems/construct-quad-tree/description/
 *
 * algorithms
 * Easy (53.41%)
 * Total Accepted:    841
 * Total Submissions: 1.6K
 * Testcase Example:  '[[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]'
 *
 * 我们想要使用一棵四叉树来储存一个 N x N 的布尔值网络。网络中每一格的值只会是真或假。树的根结点代表整个网络。对于每个结点,
 * 它将被分等成四个孩子结点直到这个区域内的值都是相同的.
 *
 * 每个结点还有另外两个布尔变量: isLeaf 和 val。isLeaf 当这个节点是一个叶子结点时为真。val 变量储存叶子结点所代表的区域的值。
 *
 * 你的任务是使用一个四叉树表示给定的网络。下面的例子将有助于你理解这个问题：
 *
 * 给定下面这个8 x 8 网络，我们将这样建立一个对应的四叉树：
 *
 *
 *
 * 由上文的定义，它能被这样分割：
 *
 *
 *
 *
 *
 * 对应的四叉树应该像下面这样，每个结点由一对 (isLeaf, val) 所代表.
 *
 * 对于非叶子结点，val 可以是任意的，所以使用 * 代替。
 *
 *
 *
 * 提示：
 *
 *
 * N 将小于 1000 且确保是 2 的整次幂。
 * 如果你想了解更多关于四叉树的知识，你可以参考这个 wiki 页面。
 *
 *
 */
/**
 * // Definition for a QuadTree node.
 * function Node(val,isLeaf,topLeft,topRight,bottomLeft,bottomRight) {
 *    this.val = val;
 *    this.isLeaf = isLeaf;
 *    this.topLeft = topLeft;
 *    this.topRight = topRight;
 *    this.bottomLeft = bottomLeft;
 *    this.bottomRight = bottomRight;
 * };
 */
/**
 * @param {number[][]} grid
 * @return {Node}
 */
var construct = function(grid) {
  return dfs(grid, 0, grid[0].length, 0, grid.length)
}

function dfs(grid, left, right, top, bottom) {
  let root = null
  let key = grid[top][left]
  for (let i = top; i < bottom; i++) {
    for (let j = left; j < right; j++) {
      if (grid[i][j] != key) {
        const topLeft = dfs(
          grid,
          left,
          ~~((left + right) / 2),
          top,
          ~~((top + bottom) / 2)
        )
        const topRight = dfs(
          grid,
          ~~((left + right) / 2),
          right,
          top,
          ~~((top + bottom) / 2)
        )
        const bottomLeft = dfs(
          grid,
          left,
          ~~((left + right) / 2),
          ~~((top + bottom) / 2),
          bottom
        )
        const bottomRight = dfs(
          grid,
          ~~((left + right) / 2),
          right,
          ~~((top + bottom) / 2),
          bottom
        )
        root = new Node(
          false,
          false,
          topLeft,
          topRight,
          bottomLeft,
          bottomRight
        )
        return root
      }
    }
  }
  root = new Node()
  root.val = key == 1 ? true : false
  root.isLeaf = true
  return root
}
