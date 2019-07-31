#
# @lc app=leetcode.cn id=179 lang=python3
#
# [179] 最大数
#
# https://leetcode-cn.com/problems/largest-number/description/
#
# algorithms
# Medium (32.25%)
# Likes:    122
# Dislikes: 0
# Total Accepted:    7.4K
# Total Submissions: 23K
# Testcase Example:  '[10,2]'
#
# 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
#
# 示例 1:
#
# 输入: [10,2]
# 输出: 210
#
# 示例 2:
#
# 输入: [3,30,34,5,9]
# 输出: 9534330
#
# 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
#
#
import functools


def compare(a, b):
    return int(str(b)+str(a)) - int(str(a)+str(b))


class Solution:

    def largestNumber(self, nums) -> str:
        res = ''
        for n in sorted(nums, key=functools.cmp_to_key(compare)):
            res += str(n)
        return '0' if res[0] == '0' else res
