#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#
# https://leetcode-cn.com/problems/permutations/description/
#
# algorithms
# Medium (67.44%)
# Total Accepted:    19.9K
# Total Submissions: 29.5K
# Testcase Example:  '[1,2,3]'
#
# 给定一个没有重复数字的序列，返回其所有可能的全排列。
# 
# 示例:
# 
# 输入: [1,2,3]
# 输出:
# [
# ⁠ [1,2,3],
# ⁠ [1,3,2],
# ⁠ [2,1,3],
# ⁠ [2,3,1],
# ⁠ [3,1,2],
# ⁠ [3,2,1]
# ]
# 
#
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if nums is None: return []
        if len(nums) == 1: return [nums]
        res = []
        for x in nums:
            ys = nums + []
            ys.remove(x)
            for y in self.permute(ys):
                res.append([x] + y)
        return res

