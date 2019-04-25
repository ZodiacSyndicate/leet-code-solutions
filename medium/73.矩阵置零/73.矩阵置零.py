#
# @lc app=leetcode.cn id=73 lang=python3
#
# [73] 矩阵置零
#


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        rowFlag = False
        for i in range(len(matrix[0])):
            if matrix[0][i] == 0:
                rowFlag = True
                break

        colFlag = False
        for i in range(len(matrix)):
            if matrix[i][0] == 0:
                colFlag = True
                break

        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[0])):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    matrix[0][j] = 0

        for i in range(1, len(matrix[0])):
            if matrix[0][i] == 0:
                for j in range(len(matrix)):
                    matrix[j][i] = 0

        for i in range(1, len(matrix)):
            if matrix[i][0] == 0:
                for j in range(len(matrix[0])):
                    matrix[i][j] = 0

        if rowFlag:
            for i in range(len(matrix[0])):
                matrix[0][i] = 0

        if colFlag:
            for i in range(len(matrix)):
                matrix[i][0] = 0
