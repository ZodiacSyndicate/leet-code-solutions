/*
 * @lc app=leetcode.cn id=200 lang=javascript
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (44.47%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    17.9K
 * Total Submissions: 40.1K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给定一个由 '1'（陆地）和
 * '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * 输出: 3
 *
 *
 */
/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
  const nr = grid.length
  if (!nr) return 0
  const nc = grid[0].length

  let res = 0
  for (let r = 0; r < nr; r++) {
    for (let c = 0; c < nc; c++) {
      if (grid[r][c] === '1') {
        res++
        dfs(grid, r, c)
      }
    }
  }
  return res
}

function dfs(grid, row, col) {
  const nr = grid.length
  const nc = grid[0].length
  grid[row][col] = '0'
  if (row - 1 >= 0 && grid[row - 1][col] === '1') dfs(grid, row - 1, col)
  if (row + 1 < nr && grid[row + 1][col] === '1') dfs(grid, row + 1, col)
  if (col - 1 >= 0 && grid[row][col - 1] === '1') dfs(grid, row, col - 1)
  if (col + 1 < nc && grid[row][col + 1] === '1') dfs(grid, row, col + 1)
}
