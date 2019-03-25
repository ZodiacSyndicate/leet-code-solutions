#
# @lc app=leetcode.cn id=33 lang=python3
#
# [33] 搜索旋转排序数组
#
# https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
#
# algorithms
# Medium (36.04%)
# Total Accepted:    18.3K
# Total Submissions: 50.8K
# Testcase Example:  '[4,5,6,7,0,1,2]\n0'
#
# 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
#
# ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
#
# 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
#
# 你可以假设数组中不存在重复的元素。
#
# 你的算法时间复杂度必须是 O(log n) 级别。
#
# 示例 1:
#
# 输入: nums = [4,5,6,7,0,1,2], target = 0
# 输出: 4
#
#
# 示例 2:
#
# 输入: nums = [4,5,6,7,0,1,2], target = 3
# 输出: -1
#
#


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.__search__(nums, target, 0, len(nums) - 1)

    def __search__(self, nums, target, low, high):
        if low > high:
            return -1
        mid = (low + high) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < nums[high]:
            if nums[mid] < target and target <= nums[high]:
                return self.__search__(nums, target, mid + 1, high)
            else:
                return self.__search__(nums, target, low, mid - 1)
        else:
            if nums[low] <= target and target < nums[mid]:
                return self.__search__(nums, target, low, mid - 1)
            else:
                return self.__search__(nums, target, mid + 1, high)
