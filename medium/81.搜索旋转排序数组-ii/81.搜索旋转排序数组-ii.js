/*
 * @lc app=leetcode.cn id=81 lang=javascript
 *
 * [81] 搜索旋转排序数组 II
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {boolean}
 */
var search = function(nums, target, i = 0, j = nums.length - 1) {
  if (j < 0) return false
  const mid = (i + j) >> 1
  if (nums[mid] === target || nums[i] === target || nums[j] === target)
    return true
  else if (mid === i) return false
  else if (nums[i] === nums[mid] && nums[mid] === nums[j])
    return search(nums, target, i, mid - 1) || search(nums, target, mid, j)
  else if (nums[mid] >= nums[i])
    return target > nums[i] && target < nums[mid]
      ? search(nums, target, i, mid - 1)
      : search(nums, target, mid + 1, j)
  else
    return target > nums[mid] && target < nums[i]
      ? search(nums, target, mid + 1, j)
      : search(nums, target, i, mid - 1)
}
