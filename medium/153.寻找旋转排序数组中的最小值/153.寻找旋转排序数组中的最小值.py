#
# @lc app=leetcode.cn id=153 lang=python3
#
# [153] 寻找旋转排序数组中的最小值
#
# https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
#
# algorithms
# Medium (49.28%)
# Likes:    60
# Dislikes: 0
# Total Accepted:    10.6K
# Total Submissions: 21.5K
# Testcase Example:  '[3,4,5,1,2]'
#
# 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
#
# ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
#
# 请找出其中最小的元素。
#
# 你可以假设数组中不存在重复元素。
#
# 示例 1:
#
# 输入: [3,4,5,1,2]
# 输出: 1
#
# 示例 2:
#
# 输入: [4,5,6,7,0,1,2]
# 输出: 0
#
#


class Solution:
    def findMin(self, nums: List[int]) -> int:
        l = 0
        r = len(nums) - 1
        if l == r:
            return nums[l]
        while l <= r:
            mid = l + ((r - l) // 2)
            if nums[l] < nums[r]:
                return nums[l]
            elif mid == 0 and nums[mid] < nums[mid + 1]:
                return nums[mid]
            elif mid == len(nums) - 1 and nums[mid] < nums[mid - 1]:
                return nums[mid]
            elif nums[mid] < nums[mid - 1] and nums[mid] < nums[mid + 1]:
                return nums[mid]
            elif nums[mid] > nums[r]:
                l = mid + 1
            elif nums[mid] < nums[l]:
                r = mid - 1
