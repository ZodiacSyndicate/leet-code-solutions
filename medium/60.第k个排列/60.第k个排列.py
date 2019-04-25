#
# @lc app=leetcode.cn id=60 lang=python3
#
# [60] 第k个排列
#
class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        arr = [x for x in range(1, n + 1)]
        res = []
        pos = k - 1
        fac = 1
        for n in arr:
            fac *= n
        for j in range(n, 0, -1):
            fac //= j
            res.append(str(arr.pop(pos // fac)))
            pos %= fac
        return ''.join(res)
