/*
 * @lc app=leetcode.cn id=59 lang=javascript
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (71.52%)
 * Total Accepted:    8.1K
 * Total Submissions: 11.2K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */
/**
 * @param {number} n
 * @return {number[][]}
 */
var generateMatrix = function(n) {
  const matrix = Array.from({ length: n }, _ =>
    Array.from({ length: n }, _ => undefined)
  )
  const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
  let x = 0,
    y = 0,
    i = 0,
    flag = false,
    k = 1
  while (true) {
    if (x < n && y < n && x >= 0 && y >= 0 && matrix[x][y] === void 0) {
      flag = false
      matrix[x][y] = k
      k++
    } else {
      if (flag) break
      flag = true
      x -= directions[i][0]
      y -= directions[i][1]
      i = (i + 1) % 4
    }
    x += directions[i][0]
    y += directions[i][1]
  }
  return matrix
}
