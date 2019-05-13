#
# @lc app=leetcode.cn id=77 lang=python3
#
# [77] 组合
#


class Solution:
    def combine(self, n: int, k: int):
        if k > n or k == 0:
            return []
        if k == 1:
            return [[x + 1] for x in range(n)]
        if k == n:
            return [[x + 1 for x in range(n)]]
        res = self.combine(n - 1, k)
        for item in self.combine(n - 1, k - 1):
            item.append(n)
            res.append(item)
        return res
