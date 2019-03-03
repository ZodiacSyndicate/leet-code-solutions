#
# @lc app=leetcode.cn id=231 lang=python3
#
# [231] 2的幂
#
# https://leetcode-cn.com/problems/power-of-two/description/
#
# algorithms
# Easy (43.83%)
# Total Accepted:    12.1K
# Total Submissions: 27.5K
# Testcase Example:  '1'
#
# 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
#
# 示例 1:
#
# 输入: 1
# 输出: true
# 解释: 20 = 1
#
# 示例 2:
#
# 输入: 16
# 输出: true
# 解释: 24 = 16
#
# 示例 3:
#
# 输入: 218
# 输出: false
#
#


class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and not (n & (n - 1))
