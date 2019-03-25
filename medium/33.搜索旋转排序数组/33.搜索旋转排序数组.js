/*
 * @lc app=leetcode.cn id=33 lang=javascript
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.04%)
 * Total Accepted:    18.3K
 * Total Submissions: 50.8K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target, low = 0, high = nums.length - 1) {
  if (low > high) return -1
  const mid = ~~((low + high) / 2)
  if (nums[mid] === target) return mid
  if (nums[mid] < nums[high]) {
    if (nums[mid] < target && target <= nums[high]) {
      return search(nums, target, mid + 1, high)
    } else {
      return search(nums, target, low, mid - 1)
    }
  } else {
    if (nums[low] <= target && target < nums[mid]) {
      return search(nums, target, low, mid - 1)
    } else {
      return search(nums, target, mid + 1, high)
    }
  }
}
