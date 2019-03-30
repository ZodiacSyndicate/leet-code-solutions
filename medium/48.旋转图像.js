/*
 * @lc app=leetcode.cn id=48 lang=javascript
 *
 * [48] 旋转图像
 *
 * https://leetcode-cn.com/problems/rotate-image/description/
 *
 * algorithms
 * Medium (60.53%)
 * Total Accepted:    19.3K
 * Total Submissions: 31.8K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵表示一个图像。
 *
 * 将图像顺时针旋转 90 度。
 *
 * 说明：
 *
 * 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
 *
 * 示例 1:
 *
 * 给定 matrix =
 * [
 * ⁠ [1,2,3],
 * ⁠ [4,5,6],
 * ⁠ [7,8,9]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [7,4,1],
 * ⁠ [8,5,2],
 * ⁠ [9,6,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 给定 matrix =
 * [
 * ⁠ [ 5, 1, 9,11],
 * ⁠ [ 2, 4, 8,10],
 * ⁠ [13, 3, 6, 7],
 * ⁠ [15,14,12,16]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [15,13, 2, 5],
 * ⁠ [14, 3, 4, 1],
 * ⁠ [12, 6, 8, 9],
 * ⁠ [16, 7,10,11]
 * ]
 *
 *
 */
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var rotate = function(matrix) {
  let n = matrix.length - 1
  let x = 0
  let y = 0
  while (n >= 0) {
    for (let i = x; i < n; i++) {
      const topLeft = matrix[y][i]
      const topRight = matrix[i][n - i]
      const bottomLeft = matrix[n - i][0]
      const bottomRight = matrix[n - i][i]
      matrix[y][i] = bottomLeft
      matrix[y][n - i] = topLeft
      matrix[n - y][i] = bottomRight
      matrix[n - y][n - i] = topRight
    }
    x++
    y++
    n -= 2
  }
}
