/*
 * @lc app=leetcode.cn id=90 lang=javascript
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (53.47%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    6.8K
 * Total Submissions: 12.5K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 */
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsetsWithDup = function(nums) {
  nums = nums.sort((a, b) => a - b)
  const res = []
  let v = []
  res.push(v)
  let [right, left, len] = [1, 0, 0]
  for (let i = 0; i < nums.length; i++) {
    if (i !== 0 && nums[i] === nums[i - 1]) {
      left = res.length - len
    } else {
      left = 0
    }
    right = res.length
    len = right - left
    for (let j = left; j < right; j++) {
      v = res[j].slice()
      v.push(nums[i])
      res.push(v)
    }
  }
  return res
}
