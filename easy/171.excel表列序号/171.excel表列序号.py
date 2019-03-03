#
# @lc app=leetcode.cn id=171 lang=python3
#
# [171] Excel表列序号
#
# https://leetcode-cn.com/problems/excel-sheet-column-number/description/
#
# algorithms
# Easy (63.07%)
# Total Accepted:    9K
# Total Submissions: 14.2K
# Testcase Example:  '"A"'
#
# 给定一个Excel表格中的列名称，返回其相应的列序号。
#
# 例如，
#
# ⁠   A -> 1
# ⁠   B -> 2
# ⁠   C -> 3
# ⁠   ...
# ⁠   Z -> 26
# ⁠   AA -> 27
# ⁠   AB -> 28
# ⁠   ...
#
#
# 示例 1:
#
# 输入: "A"
# 输出: 1
#
#
# 示例 2:
#
# 输入: "AB"
# 输出: 28
#
#
# 示例 3:
#
# 输入: "ZY"
# 输出: 701
#
# 致谢：
# 特别感谢 @ts 添加此问题并创建所有测试用例。
#
#


class Solution:
    def titleToNumber(self, s: str) -> int:
        res = 0
        for i in s:
            res = 26 * res + ord(i) - 64
        return res
