#
# @lc app=leetcode.cn id=119 lang=python3
#
# [119] 杨辉三角 II
#
# https://leetcode-cn.com/problems/pascals-triangle-ii/description/
#
# algorithms
# Easy (53.57%)
# Total Accepted:    9.7K
# Total Submissions: 18K
# Testcase Example:  '3'
#
# 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
#
#
#
# 在杨辉三角中，每个数是它左上方和右上方的数的和。
#
# 示例:
#
# 输入: 3
# 输出: [1,3,3,1]
#
#
# 进阶：
#
# 你可以优化你的算法到 O(k) 空间复杂度吗？
#
#


class Solution:
    def getRow(self, rowIndex: 'int') -> 'List[int]':
        res = [1]
        while rowIndex > 0:
            cur = []
            for i in range(len(res) + 1):
                if i == 0 or i == len(res):
                    cur.append(1)
                else:
                    cur.append(res[i] + res[i - 1])
            res = cur
            rowIndex -= 1
        return res
