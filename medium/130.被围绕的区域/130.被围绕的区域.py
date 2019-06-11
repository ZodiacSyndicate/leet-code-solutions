#
# @lc app=leetcode.cn id=130 lang=python3
#
# [130] 被围绕的区域
#
# https://leetcode-cn.com/problems/surrounded-regions/description/
#
# algorithms
# Medium (36.00%)
# Likes:    67
# Dislikes: 0
# Total Accepted:    5.3K
# Total Submissions: 14.7K
# Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
#
# 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
#
# 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
#
# 示例:
#
# X X X X
# X O O X
# X X O X
# X O X X
#
#
# 运行你的函数后，矩阵变为：
#
# X X X X
# X X X X
# X X X X
# X O X X
#
#
# 解释:
#
# 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
# 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
#
#


class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        if not board or not len(board):
            return
        m, n = len(board), len(board[0])
        for i in range(m):
            for j in range(n):
                isEdge = i == 0 or j == 0 or i == m - 1 or j == n - 1
                if isEdge and board[i][j] == 'O':
                    self.dfs(board, i, j)

        for i in range(m):
            for j in range(n):
                if board[i][j] == 'O':
                    board[i][j] = 'X'
                if board[i][j] == '#':
                    board[i][j] = 'O'

    def dfs(self, board, i, j):
        if i < 0 or j < 0 or i >= len(board) or j >= len(board[0]) or board[i][j] == 'X' or board[i][j] == '#':
            return
        board[i][j] = '#'
        self.dfs(board, i - 1, j)
        self.dfs(board, i + 1, j)
        self.dfs(board, i, j - 1)
        self.dfs(board, i, j + 1)
