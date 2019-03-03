#
# @lc app=leetcode.cn id=345 lang=python3
#
# [345] 反转字符串中的元音字母
#
# https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
#
# algorithms
# Easy (45.92%)
# Total Accepted:    6.6K
# Total Submissions: 14.4K
# Testcase Example:  '"hello"'
#
# 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
#
# 示例 1:
#
# 输入: "hello"
# 输出: "holle"
#
#
# 示例 2:
#
# 输入: "leetcode"
# 输出: "leotcede"
#
# 说明:
# 元音字母不包含字母"y"。
#
#


class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = 'aeiouAEIOU'
        s = [x for x in s]
        start = 0
        end = len(s)-1
        while start < end:
            while start < end and s[start] not in vowels:
                start += 1

            while end > start and s[end] not in vowels:
                end -= 1

            if start < end and s[start] != s[end]:
                s[start], s[end] = s[end], s[start]
            start += 1
            end -= 1
        return "".join(s)
