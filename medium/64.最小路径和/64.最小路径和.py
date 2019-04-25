#
# @lc app=leetcode.cn id=64 lang=python3
#
# [64] 最小路径和
#


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        x = len(grid)
        y = len(grid[0])
        dp = [[]] * x
        for i in range(x):
            dp[i] = [0] * y
            for j in range(y):
                if i == 0 and j == 0:
                    dp[i][j] = grid[i][j]
                elif i == 0:
                    dp[i][j] = dp[i][j - 1] + grid[i][j]
                elif j == 0:
                    dp[i][j] = dp[i - 1][j] + grid[i][j]
                else:
                    dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])
        return dp[x - 1][y - 1]
