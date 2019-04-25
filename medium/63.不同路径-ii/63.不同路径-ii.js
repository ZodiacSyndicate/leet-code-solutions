/*
 * @lc app=leetcode.cn id=63 lang=javascript
 *
 * [63] 不同路径 II
 */
/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */
var uniquePathsWithObstacles = function(obstacleGrid) {
  if (!obstacleGrid.length) return 0
  const dp = []
  const x = obstacleGrid.length
  const y = obstacleGrid[0].length
  for (let i = 0; i < x; i++) {
    dp[i] = []
    for (let j = 0; j < y; j++) {
      if (obstacleGrid[i][j] === 1) {
        dp[i][j] = 0
      } else {
        if (i === 0 && j === 0) {
          dp[i][j] = 1
        } else if (i === 0) {
          dp[i][j] = dp[i][j - 1]
        } else if (j === 0) {
          dp[i][j] = dp[i - 1][j]
        } else {
          dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        }
      }
    }
  }
  return dp[x - 1][y - 1]
}
