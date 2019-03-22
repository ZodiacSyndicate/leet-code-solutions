#
# @lc app=leetcode.cn id=18 lang=python3
#
# [18] 四数之和
#
# https://leetcode-cn.com/problems/4sum/description/
#
# algorithms
# Medium (34.65%)
# Total Accepted:    14K
# Total Submissions: 40.3K
# Testcase Example:  '[1,0,-1,0,-2,2]\n0'
#
# 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
# + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
#
# 注意：
#
# 答案中不可以包含重复的四元组。
#
# 示例：
#
# 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
#
# 满足要求的四元组集合为：
# [
# ⁠ [-1,  0, 0, 1],
# ⁠ [-2, -1, 1, 2],
# ⁠ [-2,  0, 0, 2]
# ]
#
#
#


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums = sorted(nums)
        return self.Sum(nums, target, 4)

    def Sum(self, nums, target, flag):
        res = []
        if flag == 1:
            if target in nums:
                return [[target]]
            return
        else:
            for i in range(len(nums)):
                if i != 0 and nums[i] == nums[i - 1]:
                    continue
                temp = self.Sum(nums[i+1:], target - nums[i], flag - 1)
                if temp:
                    for arr in temp:
                        arr.append(nums[i])
                        res.append(arr)
        return res
