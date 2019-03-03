#
# @lc app=leetcode.cn id=118 lang=python3
#
# [118] 杨辉三角
#
# https://leetcode-cn.com/problems/pascals-triangle/description/
#
# algorithms
# Easy (59.45%)
# Total Accepted:    15.9K
# Total Submissions: 26.6K
# Testcase Example:  '5'
#
# 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
#
#
#
# 在杨辉三角中，每个数是它左上方和右上方的数的和。
#
# 示例:
#
# 输入: 5
# 输出:
# [
# ⁠    [1],
# ⁠   [1,1],
# ⁠  [1,2,1],
# ⁠ [1,3,3,1],
# ⁠[1,4,6,4,1]
# ]
#
#


class Solution:
    def generate(self, numRows: 'int') -> 'List[List[int]]':
        if not numRows:
            return []
        res = [[1]]
        arr = [1]
        while numRows > 1:
            cur = []
            for j in range(len(arr) + 1):
                if j == 0 or j == len(arr):
                    cur.append(1)
                else:
                    cur.append(arr[j] + arr[j - 1])
            res.append(cur)
            arr = cur
            numRows -= 1
        return res
