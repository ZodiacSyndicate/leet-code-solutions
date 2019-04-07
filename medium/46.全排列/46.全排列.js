/*
 * @lc app=leetcode.cn id=46 lang=javascript
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (67.08%)
 * Total Accepted:    18.7K
 * Total Submissions: 27.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function(nums, start = 0, end = nums.length, result = []) {
  if (start === end) {
    result.push(nums)
  } else {
    for (let i = start; i < end; i++) {
      ;[nums[i], nums[start]] = [nums[start], nums[i]]
      permute([...nums], start + 1, end, result)
      ;[nums[i], nums[start]] = [nums[start], nums[i]]
    }
  }
  return result
}
