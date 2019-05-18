#
# @lc app=leetcode.cn id=81 lang=python3
#
# [81] 搜索旋转排序数组 II
#
class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        return self.mySearch(nums, target, 0, len(nums) - 1)

    def mySearch(self, nums, target, i, j):
        if j < 0:
            return False
        mid = (i + j) >> 1
        if nums[mid] == target or nums[i] == target or nums[j] == target:
            return True
        elif mid == i:
            return False
        elif nums[i] == nums[mid] and nums[mid] == nums[j]:
            return self.mySearch(nums, target, i, mid-1) or self.mySearch(nums, target, mid, j)
        elif nums[mid] >= nums[i]:
            if target > nums[i] and target < nums[mid]:
                return self.mySearch(nums, target, i, mid - 1)
            else:
                return self.mySearch(nums, target, mid + 1, j)
        else:
            if target > nums[mid] and target < nums[i]:
                return self.mySearch(nums, target, mid + 1, j) 
            else:
                return self.mySearch(nums, target, i, mid - 1)
