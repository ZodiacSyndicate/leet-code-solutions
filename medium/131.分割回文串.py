#
# @lc app=leetcode.cn id=131 lang=python3
#
# [131] 分割回文串
#
# https://leetcode-cn.com/problems/palindrome-partitioning/description/
#
# algorithms
# Medium (62.32%)
# Likes:    115
# Dislikes: 0
# Total Accepted:    7.7K
# Total Submissions: 12.3K
# Testcase Example:  '"aab"'
#
# 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
#
# 返回 s 所有可能的分割方案。
#
# 示例:
#
# 输入: "aab"
# 输出:
# [
# ⁠ ["aa","b"],
# ⁠ ["a","a","b"]
# ]
#
#


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        return self.backtrack(s, [], [])

    def backtrack(self, str, arr, res):
        if self.isPalidrome(str):
            temp = arr[:]
            temp.append(str)
            res.append(temp)
        for i in range(1, len(str)):
            temp = arr[:]
            s = str[0:i]
            if not self.isPalidrome(s):
                continue
            temp.append(s)
            self.backtrack(str[i:], temp, res)
        return res

    def isPalidrome(self, s):
        if len(s) == 1:
            return True
        i, j = 0, len(s) - 1
        while i <= j:
            if s[i] != s[j]:
                return False
            i += 1
            j -= 1
        return True
