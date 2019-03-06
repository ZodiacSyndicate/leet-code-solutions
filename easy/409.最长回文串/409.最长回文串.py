#
# @lc app=leetcode.cn id=409 lang=python3
#
# [409] 最长回文串
#
# https://leetcode-cn.com/problems/longest-palindrome/description/
#
# algorithms
# Easy (45.63%)
# Total Accepted:    3.6K
# Total Submissions: 7.8K
# Testcase Example:  '"abccccdd"'
#
# 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
#
# 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
#
# 注意:
# 假设字符串的长度不会超过 1010。
#
# 示例 1:
#
#
# 输入:
# "abccccdd"
#
# 输出:
# 7
#
# 解释:
# 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
#
#
#


class Solution:
    def longestPalindrome(self, s: str) -> int:
        m = {}
        ans = 0
        flag = False
        for n in s:
            if n in m:
                m[n] += 1
            else:
                m[n] = 1
        for n in m:
            if m[n] % 2 == 0:
                ans += m[n]
            else:
                flag = True
                ans += m[n] - 1
        return ans + 1 if flag else ans
