#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#
# https://leetcode-cn.com/problems/3sum/description/
#
# algorithms
# Medium (21.45%)
# Total Accepted:    41.5K
# Total Submissions: 193.6K
# Testcase Example:  '[-1,0,1,2,-1,-4]'
#
# 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
# ？找出所有满足条件且不重复的三元组。
#
# 注意：答案中不可以包含重复的三元组。
#
# 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
#
# 满足要求的三元组集合为：
# [
# ⁠ [-1, 0, 1],
# ⁠ [-1, -1, 2]
# ]
#
#
#


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        ans = []
        c = len(nums) - 1
        while c >= 2:
            a, b = 0, c - 1
            while a < b:
                s = nums[a] + nums[b]
                if s < -nums[c]:
                    a += 1
                elif s > - nums[c]:
                    b -= 1
                else:
                    ans.append([nums[a], nums[b], nums[c]])
                    a += 1
                    while a < b and nums[a - 1] == nums[a]:
                        a += 1
                    b -= 1
                    while a < b and nums[b + 1] == nums[b]:
                        b -= 1
            c -= 1
            while c >= 2 and nums[c + 1] == nums[c]:
                c -= 1
        return ans
