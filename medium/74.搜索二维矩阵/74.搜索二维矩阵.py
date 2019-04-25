#
# @lc app=leetcode.cn id=74 lang=python3
#
# [74] 搜索二维矩阵
#


class Solution:
    def searchMatrix(self, matrix, target: int) -> bool:
        if len(matrix) == 0:
            return False
        x, y = len(matrix), len(matrix[0])
        l, r = 0, x * y - 1
        while l <= r:
            mid = (l + r) // 2
            ix, iy = mid // y, mid % y
            if matrix[ix][iy] == target:
                return True
            elif matrix[ix][iy] > target:
                r = mid - 1
            else:
                l = mid + 1
        return False
