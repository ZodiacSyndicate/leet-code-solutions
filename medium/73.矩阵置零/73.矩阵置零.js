/*
 * @lc app=leetcode.cn id=73 lang=javascript
 *
 * [73] 矩阵置零
 */
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var setZeroes = function(matrix) {
  const xSet = new Set()
  const ySet = new Set()
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (matrix[i][j] === 0) {
        xSet.add(i)
        ySet.add(j)
      }
    }
  }
  for (let n of xSet.values()) {
    matrix[n].fill(0)
  }
  for (let n of ySet.values()) {
    for (let i = 0; i < matrix.length; i++) {
      matrix[i][n] = 0
    }
  }
}
