#
# @lc app=leetcode.cn id=447 lang=python3
#
# [447] 回旋镖的数量
#
# https://leetcode-cn.com/problems/number-of-boomerangs/description/
#
# algorithms
# Easy (52.92%)
# Total Accepted:    3.2K
# Total Submissions: 5.9K
# Testcase Example:  '[[0,0],[1,0],[2,0]]'
#
# 给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k
# 之间的距离相等（需要考虑元组的顺序）。
#
# 找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
#
# 示例:
#
#
# 输入:
# [[0,0],[1,0],[2,0]]
#
# 输出:
# 2
#
# 解释:
# 两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
#
#
#


class Solution:
    def numberOfBoomerangs(self, points: List[List[int]]) -> int:
        ans = 0
        for i in range(len(points)):
            m = {}
            for j in range(len(points)):
                a1, b1 = points[i]
                a2, b2 = points[j]
                dis = (a1 - a2) ** 2 + (b1 - b2) ** 2
                if dis in m:
                    m[dis] += 1
                else:
                    m[dis] = 1
            for n in m.values():
                if n > 1:
                    ans += n * (n - 1)
        return ans
