#
# @lc app=leetcode.cn id=387 lang=python3
#
# [387] 字符串中的第一个唯一字符
#
# https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
#
# algorithms
# Easy (35.49%)
# Total Accepted:    21.3K
# Total Submissions: 59.4K
# Testcase Example:  '"leetcode"'
#
# 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
#
# 案例:
#
#
# s = "leetcode"
# 返回 0.
#
# s = "loveleetcode",
# 返回 2.
#
#
#
#
# 注意事项：您可以假定该字符串只包含小写字母。
#
#


class Solution:
    def firstUniqChar(self, s: str) -> int:
        arr = [0] * 26
        for n in s:
            arr[ord(n) - ord('a')] += 1
        for i in range(len(s)):
            if arr[ord(s[i]) - ord('a')] == 1:
                return i
        return -1
