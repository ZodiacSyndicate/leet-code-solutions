#
# @lc app=leetcode.cn id=459 lang=python3
#
# [459] 重复的子字符串
#
# https://leetcode-cn.com/problems/repeated-substring-pattern/description/
#
# algorithms
# Easy (38.79%)
# Total Accepted:    4K
# Total Submissions: 10.2K
# Testcase Example:  '"abab"'
#
# 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
#
# 示例 1:
#
#
# 输入: "abab"
#
# 输出: True
#
# 解释: 可由子字符串 "ab" 重复两次构成。
#
#
# 示例 2:
#
#
# 输入: "aba"
#
# 输出: False
#
#
# 示例 3:
#
#
# 输入: "abcabcabcabc"
#
# 输出: True
#
# 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
#
#
#


class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        arr = []
        prev = 0
        for i in range(1, len(s)):
            if s[i] == s[0]:
                arr.append(i)
            if len(s) % (i - prev) != 0:
                return False
            prev = i
        if not len(arr):
            return False
        for idx in arr:
            if s[0:idx] * (len(s) // idx) == s:
                return True
        return False
