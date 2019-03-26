#
# @lc app=leetcode.cn id=34 lang=python3
#
# [34] 在排序数组中查找元素的第一个和最后一个位置
#
# https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
#
# algorithms
# Medium (35.44%)
# Total Accepted:    13.8K
# Total Submissions: 38.8K
# Testcase Example:  '[5,7,7,8,8,10]\n8'
#
# 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
#
# 你的算法时间复杂度必须是 O(log n) 级别。
#
# 如果数组中不存在目标值，返回 [-1, -1]。
#
# 示例 1:
#
# 输入: nums = [5,7,7,8,8,10], target = 8
# 输出: [3,4]
#
# 示例 2:
#
# 输入: nums = [5,7,7,8,8,10], target = 6
# 输出: [-1,-1]
#
#


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        l, r = 0, len(nums) - 1
        ans = []
        while l <= r:
            mid = l + (r - l) // 2
            if nums[mid] == target:
                for i in range(l, r + 1):
                    if nums[i] == target:
                        ans.append(i)
                break
            elif nums[mid] < target:
                l = mid + 1
            else:
                r = mid - 1
        return [-1, -1] if len(ans) == 0 else [ans[0], ans[-1]]
