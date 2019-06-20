#
# @lc app=leetcode.cn id=151 lang=python3
#
# [151] 翻转字符串里的单词
#
# https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
#
# algorithms
# Medium (25.76%)
# Likes:    56
# Dislikes: 0
# Total Accepted:    11.5K
# Total Submissions: 43.7K
# Testcase Example:  '"the sky is blue"'
#
# 给定一个字符串，逐个翻转字符串中的每个单词。
#
#
#
# 示例 1：
#
# 输入: "the sky is blue"
# 输出: "blue is sky the"
#
#
# 示例 2：
#
# 输入: "  hello world!  "
# 输出: "world! hello"
# 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
#
#
# 示例 3：
#
# 输入: "a good   example"
# 输出: "example good a"
# 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
#
#
#
#
# 说明：
#
#
# 无空格字符构成一个单词。
# 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
# 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
#
#
#
#
# 进阶：
#
# 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
#
#


class Solution:
    def reverseWords(self, s: str) -> str:
        s = s.strip()
        cur = ''
        arr = []
        for n in s:
            if n != ' ':
                cur += n
            else:
                if cur:
                    arr.append(cur)
                    cur = ''
        arr.append(cur)
        arr.reverse()
        return ' '.join(arr)
