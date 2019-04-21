#
# @lc app=leetcode.cn id=59 lang=python3
#
# [59] 螺旋矩阵 II
#
# https://leetcode-cn.com/problems/spiral-matrix-ii/description/
#
# algorithms
# Medium (71.52%)
# Total Accepted:    8.1K
# Total Submissions: 11.2K
# Testcase Example:  '3'
#
# 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
#
# 示例:
#
# 输入: 3
# 输出:
# [
# ⁠[ 1, 2, 3 ],
# ⁠[ 8, 9, 4 ],
# ⁠[ 7, 6, 5 ]
# ]
#
#


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        matrix = []
        for i in range(n):
            l = []
            for j in range(n):
                l.append(None)
            matrix.append(l)
        directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
        x, y, i, k, flag = 0, 0, 0, 1, False
        while True:
            if x < n and y < n and x >= 0 and y >= 0 and matrix[x][y] == None:
                flag = False
                matrix[x][y] = k
                k += 1
            else:
                if flag:
                    break
                flag = True
                x -= directions[i][0]
                y -= directions[i][1]
                i = (i + 1) % 4
            x += directions[i][0]
            y += directions[i][1]
        return matrix
