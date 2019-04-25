#
# @lc app=leetcode.cn id=63 lang=python3
#
# [63] 不同路径 II
#


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        if len(obstacleGrid) == 0:
            return 0
        x = len(obstacleGrid)
        y = len(obstacleGrid[0])
        dp = [[]] * x
        for i in range(x):
            dp[i] = [0] * y
            for j in range(y):
                if obstacleGrid[i][j] == 1:
                    dp[i][j] = 0
                else:
                    if i == 0 and j == 0:
                        dp[i][j] = 1
                    elif i == 0:
                        dp[i][j] = dp[i][j - 1]
                    elif j == 0:
                        dp[i][j] = dp[i - 1][j]
                    else:
                        dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        return dp[x - 1][y - 1]
