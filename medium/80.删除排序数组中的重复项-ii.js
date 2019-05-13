/*
 * @lc app=leetcode.cn id=80 lang=javascript
 *
 * [80] 删除排序数组中的重复项 II
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  let i = 0
  for (const n of nums) {
    if (i < 2 || n > nums[i - 2]) {
      nums[i++] = n
    }
  }
  return i
}
