#
# @lc app=leetcode.cn id=69 lang=python3
#
# [69] x 的平方根
#
# https://leetcode-cn.com/problems/sqrtx/description/
#
# algorithms
# Easy (34.28%)
# Total Accepted:    22.7K
# Total Submissions: 65.9K
# Testcase Example:  '4'
#
# 实现 int sqrt(int x) 函数。
#
# 计算并返回 x 的平方根，其中 x 是非负整数。
#
# 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
#
# 示例 1:
#
# 输入: 4
# 输出: 2
#
#
# 示例 2:
#
# 输入: 8
# 输出: 2
# 说明: 8 的平方根是 2.82842...,
# 由于返回类型是整数，小数部分将被舍去。
#
#
#
class Solution:
    def mySqrt(self, x: 'int') -> 'int':
        l = 0
        r = x
        while l <= x:
            mid = l + (r - l) // 2
            if mid ** 2 <= x and (mid + 1) ** 2 > x:
                return mid
            elif mid ** 2 > x:
                r = mid - 1
            else:
                l = mid + 1
