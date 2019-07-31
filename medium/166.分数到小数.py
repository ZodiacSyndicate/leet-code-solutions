#
# @lc app=leetcode.cn id=166 lang=python3
#
# [166] 分数到小数
#
# https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
#
# algorithms
# Medium (24.01%)
# Likes:    41
# Dislikes: 0
# Total Accepted:    3.2K
# Total Submissions: 13.4K
# Testcase Example:  '1\n2'
#
# 给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
#
# 如果小数部分为循环小数，则将循环的部分括在括号内。
#
# 示例 1:
#
# 输入: numerator = 1, denominator = 2
# 输出: "0.5"
#
#
# 示例 2:
#
# 输入: numerator = 2, denominator = 1
# 输出: "2"
#
# 示例 3:
#
# 输入: numerator = 2, denominator = 3
# 输出: "0.(6)"
#
#
#


class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        if numerator == 0:
            return '0'
        fraction = ''
        if (numerator < 0 and denominator > 0) or (numerator > 0 and denominator < 0):
            fraction += '-'
        dividend = abs(numerator)
        divisor = abs(denominator)
        fraction += str(dividend // divisor)
        remainder = dividend % divisor
        if remainder == 0:
            return fraction
        fraction += '.'
        m = {}
        while remainder != 0:
            if m.get(remainder):
                fraction = fraction[0:m[remainder]] + \
                    '(' + fraction[m[remainder]:] + ')'
                break
            m[remainder] = len(fraction)
            remainder *= 10
            fraction += str(remainder // divisor)
            remainder %= divisor
        return fraction
