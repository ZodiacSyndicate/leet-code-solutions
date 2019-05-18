#
# @lc app=leetcode.cn id=90 lang=python3
#
# [90] 子集 II
#
# https://leetcode-cn.com/problems/subsets-ii/description/
#
# algorithms
# Medium (53.16%)
# Likes:    85
# Dislikes: 0
# Total Accepted:    6.9K
# Total Submissions: 12.8K
# Testcase Example:  '[1,2,2]'
#
# 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
#
# 说明：解集不能包含重复的子集。
#
# 示例:
#
# 输入: [1,2,2]
# 输出:
# [
# ⁠ [2],
# ⁠ [1],
# ⁠ [1,2,2],
# ⁠ [2,2],
# ⁠ [1,2],
# ⁠ []
# ]
#
#


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        res = []
        v = []
        res.append(v)
        right, left, l = 1, 0, 0
        for i in range(len(nums)):
            if i != 0 and nums[i] == nums[i - 1]:
                left = len(res) - l
            else:
                left = 0
            right = len(res)
            l = right - left
            for j in range(left, right):
                v = res[j].copy()
                v.append(nums[i])
                res.append(v)
        return res
