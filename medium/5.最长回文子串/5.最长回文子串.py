#
# @lc app=leetcode.cn id=5 lang=python3
#
# [5] 最长回文子串
#
# https://leetcode-cn.com/problems/longest-palindromic-substring/description/
#
# algorithms
# Medium (24.64%)
# Total Accepted:    40.1K
# Total Submissions: 162.6K
# Testcase Example:  '"babad"'
#
# 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
#
# 示例 1：
#
# 输入: "babad"
# 输出: "bab"
# 注意: "aba" 也是一个有效答案。
#
#
# 示例 2：
#
# 输入: "cbbd"
# 输出: "bb"
#
#
#


class Solution:
    def longestPalindrome(self, s: str) -> str:
        arr = ['$', '#']
        for n in s:
            arr.append(n)
            arr.append('#')
        arr.append('%')

        ans = ''
        maxPoint = 0
        rightMax = 0
        middle = 0
        lens = [1] * len(arr)
        for i in range(len(arr)):
            if len(arr) - 1 - i <= lens[maxPoint]:
                break

            if rightMax > i:
                lens[i] = min(rightMax - i, lens[2 * middle - i])

            while arr[i - lens[i]] == arr[i + lens[i]]:
                lens[i] += 1

            if lens[i] + i > rightMax:
                middle = i
                rightMax = lens[i] + i

            if lens[i] > lens[maxPoint]:
                maxPoint = i

        for i in range(maxPoint - (lens[maxPoint] - 2), maxPoint + lens[maxPoint] - 1, 2):
            ans += arr[i]

        return ans
