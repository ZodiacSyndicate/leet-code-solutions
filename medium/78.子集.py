#
# @lc app=leetcode.cn id=78 lang=python3
#
# [78] 子集
#


class Solution:
    def subsets(self, nums):
        res = [[]]
        for n in nums:
            l = len(res)
            for i in range(l):
                res.append(res[i] + [n])
        return res
