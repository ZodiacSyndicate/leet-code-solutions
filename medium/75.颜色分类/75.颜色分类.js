/*
 * @lc app=leetcode.cn id=75 lang=javascript
 *
 * [75] 颜色分类
 */
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
  let [i, j, k] = [0, 0, nums.length - 1]
  while (j <= k) {
    if (nums[j] === 1) {
      j++
    } else if (nums[j] === 0) {
      ;[nums[i], nums[j]] = [nums[j], nums[i]]
      i++
      j++
    } else {
      ;[nums[j], nums[k]] = [nums[k], nums[j]]
      k--
    }
  }
}
