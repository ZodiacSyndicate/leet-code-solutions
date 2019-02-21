#
# @lc app=leetcode.cn id=7 lang=python3
#
# [7] 整数反转
#
# https://leetcode-cn.com/problems/reverse-integer/description/
#
# algorithms
# Easy (31.41%)
# Total Accepted:    79.4K
# Total Submissions: 252.8K
# Testcase Example:  '123'
#
# 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
#
# 示例 1:
#
# 输入: 123
# 输出: 321
#
#
# 示例 2:
#
# 输入: -123
# 输出: -321
#
#
# 示例 3:
#
# 输入: 120
# 输出: 21
#
#
# 注意:
#
# 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
#
#
class Solution:
    def reverse(self, x: 'int') -> 'int':
        flag = False
        if x < 0:
            flag = True
        res = 0
        if x < 0:
            x = - x
        while x != 0:
            n = x % 10
            res = res * 10 + n
            if res < -(2 ** 31) or res > 2 ** 31 - 1:
                return 0
            x //= 10
        if flag:
            res = -res
        return res
