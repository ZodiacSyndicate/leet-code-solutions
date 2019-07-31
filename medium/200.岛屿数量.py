#
# @lc app=leetcode.cn id=200 lang=python3
#
# [200] 岛屿数量
#
# https://leetcode-cn.com/problems/number-of-islands/description/
#
# algorithms
# Medium (44.47%)
# Likes:    180
# Dislikes: 0
# Total Accepted:    17.9K
# Total Submissions: 40.1K
# Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
#
# 给定一个由 '1'（陆地）和
# '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
#
# 示例 1:
#
# 输入:
# 11110
# 11010
# 11000
# 00000
#
# 输出: 1
#
#
# 示例 2:
#
# 输入:
# 11000
# 11000
# 00100
# 00011
#
# 输出: 3
#
#
#


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        nr = len(grid)
        if not nr:
            return 0
        nc = len(grid[0])

        res = 0
        for r in range(nr):
            for c in range(nc):
                if grid[r][c] == '1':
                    res += 1
                    self.dfs(grid, r, c)

        return res

    def dfs(self, grid, row, col):
        nr = len(grid)
        nc = len(grid[0])
        grid[row][col] = '0'
        if row - 1 >= 0 and grid[row - 1][col] == '1':
            self.dfs(grid, row - 1, col)
        if row + 1 < nr and grid[row + 1][col] == '1':
            self.dfs(grid, row + 1, col)
        if col - 1 >= 0 and grid[row][col - 1] == '1':
            self.dfs(grid, row, col - 1)
        if col + 1 < nc and grid[row][col + 1] == '1':
            self.dfs(grid, row, col + 1)
