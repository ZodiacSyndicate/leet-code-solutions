#
# @lc app=leetcode.cn id=168 lang=python3
#
# [168] Excel表列名称
#
# https://leetcode-cn.com/problems/excel-sheet-column-title/description/
#
# algorithms
# Easy (29.78%)
# Total Accepted:    5.1K
# Total Submissions: 17K
# Testcase Example:  '1'
#
# 给定一个正整数，返回它在 Excel 表中相对应的列名称。
#
# 例如，
#
# ⁠   1 -> A
# ⁠   2 -> B
# ⁠   3 -> C
# ⁠   ...
# ⁠   26 -> Z
# ⁠   27 -> AA
# ⁠   28 -> AB
# ⁠   ...
#
#
# 示例 1:
#
# 输入: 1
# 输出: "A"
#
#
# 示例 2:
#
# 输入: 28
# 输出: "AB"
#
#
# 示例 3:
#
# 输入: 701
# 输出: "ZY"
#
#
#


class Solution:
    def convertToTitle(self, n: int) -> str:
        arr = ['Z']
        res = ''
        for i in range(ord('A'), ord('Z')):
            arr.append(chr(i))
        while n != 0:
            num = n % 26
            res = str(arr[num]) + res
            n = (n - num) // 26
            if num == 0:
                n -= 1
        return res
