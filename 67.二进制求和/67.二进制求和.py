#
# @lc app=leetcode.cn id=67 lang=python3
#
# [67] 二进制求和
#
# https://leetcode-cn.com/problems/add-binary/description/
#
# algorithms
# Easy (46.14%)
# Total Accepted:    15.9K
# Total Submissions: 34.3K
# Testcase Example:  '"11"\n"1"'
#
# 给定两个二进制字符串，返回他们的和（用二进制表示）。
#
# 输入为非空字符串且只包含数字 1 和 0。
#
# 示例 1:
#
# 输入: a = "11", b = "1"
# 输出: "100"
#
# 示例 2:
#
# 输入: a = "1010", b = "1011"
# 输出: "10101"
#
#
class Solution:
    def addBinary(self, a: 'str', b: 'str') -> 'str':
        res = ''
        flag = False
        while a or b:
            if not a:
                a = '0'
            if not b:
                b = '0'
            a1, b1 = int(a[-1]), int(b[-1])
            if flag:
                if a1 + b1 + 1 < 2:
                    flag = False
                res = str((a1 + b1 + 1) % 2) + res
            else:
                if a1 + b1 == 2:
                    flag = True
                res = str((a1 + b1) % 2) + res
            a, b = a[0: len(a) - 1], b[0: len(b) - 1]
        if flag:
            res = '1' + res
        return res
