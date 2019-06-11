/*
 * @lc app=leetcode.cn id=130 lang=javascript
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (36.00%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    5.3K
 * Total Submissions: 14.7K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 *
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 *
 * 示例:
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * 运行你的函数后，矩阵变为：
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * 解释:
 *
 * 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 */
/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solve = function(board) {
  if (!board || !board.length) return
  let m = board.length
  let n = board[0].length
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      const isEdge = i === 0 || j === 0 || i === m - 1 || j === n - 1
      if (isEdge && board[i][j] === 'O') {
        dfs(board, i, j)
      }
    }
  }

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (board[i][j] === 'O') {
        board[i][j] = 'X'
      }
      if (board[i][j] === '#') {
        board[i][j] = 'O'
      }
    }
  }
}

function dfs(board, i, j) {
  if (
    i < 0 ||
    j < 0 ||
    i >= board.length ||
    j >= board[0].length ||
    board[i][j] == 'X' ||
    board[i][j] == '#'
  ) {
    return
  }
  board[i][j] = '#'
  dfs(board, i - 1, j)
  dfs(board, i + 1, j)
  dfs(board, i, j - 1)
  dfs(board, i, j + 1)
}
