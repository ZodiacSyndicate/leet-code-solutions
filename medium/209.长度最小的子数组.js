/*
 * @lc app=leetcode.cn id=209 lang=javascript
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (38.97%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 30.1K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例:
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 */
/**
 * @param {number} s
 * @param {number[]} nums
 * @return {number}
 */
var minSubArrayLen = function(s, nums) {
  if (nums.length === 0) return 0
  let l = 0
  let r = 0
  let len = Number.MAX_SAFE_INTEGER
  let sum = 0
  while (r <= nums.length) {
    if (sum >= s) {
      if (r - l < len) {
        len = Math.min(len, r - l)
      }
      sum -= nums[l]
      l++
    } else {
      sum += nums[r]
      r++
    }
  }
  return len == Number.MAX_SAFE_INTEGER ? 0 : len
}
