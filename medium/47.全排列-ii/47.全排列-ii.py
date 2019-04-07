#
# @lc app=leetcode.cn id=47 lang=python3
#
# [47] 全排列 II
#
# https://leetcode-cn.com/problems/permutations-ii/description/
#
# algorithms
# Medium (51.19%)
# Total Accepted:    9.4K
# Total Submissions: 18.4K
# Testcase Example:  '[1,1,2]'
#
# 给定一个可包含重复数字的序列，返回所有不重复的全排列。
# 
# 示例:
# 
# 输入: [1,1,2]
# 输出:
# [
# ⁠ [1,1,2],
# ⁠ [1,2,1],
# ⁠ [2,1,1]
# ]
# 
#
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        ans = []
        self.perm(nums, 0, len(nums) - 1, ans)
        return ans
     
    def perm(self, nums, left, right, ans):
        if left == right:
            ans.append(nums)
        else:
            for i in range(left, right + 1):
                if i != left and nums[left] == nums[i]:
                    continue
                nums[left], nums[i] = nums[i], nums[left]
                self.perm(nums[:], left + 1, right, ans) 

