#
# @lc app=leetcode.cn id=54 lang=python3
#
# [54] 螺旋矩阵
#
# https://leetcode-cn.com/problems/spiral-matrix/description/
#
# algorithms
# Medium (34.30%)
# Total Accepted:    11.4K
# Total Submissions: 32.9K
# Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
#
# 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
#
# 示例 1:
#
# 输入:
# [
# ⁠[ 1, 2, 3 ],
# ⁠[ 4, 5, 6 ],
# ⁠[ 7, 8, 9 ]
# ]
# 输出: [1,2,3,6,9,8,7,4,5]
#
#
# 示例 2:
#
# 输入:
# [
# ⁠ [1, 2, 3, 4],
# ⁠ [5, 6, 7, 8],
# ⁠ [9,10,11,12]
# ]
# 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
#
#
#


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
        res = []
        x, y, i, flag = 0, 0, 0, False
        while True:
            if x < len(matrix) and y < len(matrix[x]) and matrix[x][y] != None:
                flag = False
                res.append(matrix[x][y])
                matrix[x][y] = None
            else:
                if flag:
                    break
                flag = True
                x -= directions[i][0]
                y -= directions[i][1]
                i = (i + 1) % 4
            x += directions[i][0]
            y += directions[i][1]
        return res
