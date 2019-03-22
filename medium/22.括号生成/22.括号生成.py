#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
#
# https://leetcode-cn.com/problems/generate-parentheses/description/
#
# algorithms
# Medium (68.93%)
# Total Accepted:    16.8K
# Total Submissions: 24.3K
# Testcase Example:  '3'
#
# 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
#
# 例如，给出 n = 3，生成结果为：
#
# [
# ⁠ "((()))",
# ⁠ "(()())",
# ⁠ "(())()",
# ⁠ "()(())",
# ⁠ "()()()"
# ]
#
#
#


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        ans = []

        def backtrack(s='', l=0, r=0):
            if len(s) == 2 * n:
                ans.append(s)
                return
            if l < n:
                backtrack(s + '(', l + 1, r)
            if r < l:
                backtrack(s + ')', l, r + 1)
        backtrack()
        return ans
