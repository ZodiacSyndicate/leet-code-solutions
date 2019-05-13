#
# @lc app=leetcode.cn id=79 lang=python3
#
# [79] 单词搜索
#


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        for i in range(len(board)):
            for j in range(len(board[i])):
                if self.dfs(board, i, j, word, 0):
                    return True
        return False

    def dfs(self, board, i, j, word, k):
        if k == len(word):
            return True
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != word[k]:
            return False
        s = board[i][j]
        board[i][j] = '*'
        res = self.dfs(board, i+1, j, word, k+1) or self.dfs(board, i-1, j, word, k +
                                                             1)or self.dfs(board, i, j+1, word, k+1)or self.dfs(board, i, j-1, word, k+1)
        board[i][j] = s
        return res
