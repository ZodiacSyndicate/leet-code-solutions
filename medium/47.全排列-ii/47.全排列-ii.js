/*
 * @lc app=leetcode.cn id=47 lang=javascript
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (51.10%)
 * Total Accepted:    8.8K
 * Total Submissions: 17.2K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permuteUnique = function(nums) {
  nums = nums.sort((a, b) => a - b)
  const ans = []
  perm(nums, 0, nums.length - 1, ans)
  return ans
}

function perm(nums, left, right, ans) {
  if (left === right) {
    ans.push(nums)
  } else {
    for (let i = left; i <= right; i++) {
      if (i !== left && nums[left] === nums[i]) continue
      ;[nums[left], nums[i]] = [nums[i], nums[left]]
      perm([...nums], left + 1, right, ans)
    }
  }
}
