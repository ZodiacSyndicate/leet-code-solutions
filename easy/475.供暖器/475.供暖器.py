#
# @lc app=leetcode.cn id=475 lang=python3
#
# [475] 供暖器
#
# https://leetcode-cn.com/problems/heaters/description/
#
# algorithms
# Easy (26.72%)
# Total Accepted:    1.9K
# Total Submissions: 7.1K
# Testcase Example:  '[1,2,3]\n[2]'
#
# 冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
#
# 现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
#
# 所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。
#
# 说明:
#
#
# 给出的房屋和供暖器的数目是非负数且不会超过 25000。
# 给出的房屋和供暖器的位置均是非负数且不会超过10^9。
# 只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
# 所有供暖器都遵循你的半径标准，加热的半径也一样。
#
#
# 示例 1:
#
#
# 输入: [1,2,3],[2]
# 输出: 1
# 解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
#
#
# 示例 2:
#
#
# 输入: [1,2,3,4],[1,4]
# 输出: 1
# 解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
#
#
#


class Solution:
    def findRadius(self, houses: List[int], heaters: List[int]) -> int:
        houses = sorted(houses)
        heaters = sorted(heaters)
        l1, l2 = len(heaters), len(houses)
        if l1 == 1:
            return max(abs(houses[0] - heaters[0]), abs(houses[-1] - heaters[0]))
        res = 0
        j = 0
        m = l2 - 1
        while m >= 0 and houses[m] >= heaters[-1]:
            m -= 1
        if m != l2 - 1:
            res = houses[-1] - heaters[-1]
        n = 0
        while n < l2 and houses[n] <= heaters[0]:
            n += 1
        if n != 0:
            res = max(res, heaters[0] - houses[0])
        for i in range(n, m + 1):
            while houses[i] > heaters[j]:
                j += 1
            if houses[i] == heaters[j]:
                continue
            res = max(
                res, min(houses[i] - heaters[j - 1], heaters[j] - houses[i]))
        return res
