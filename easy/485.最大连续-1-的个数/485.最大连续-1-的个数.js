/*
 * @lc app=leetcode.cn id=485 lang=javascript
 *
 * [485] 最大连续1的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (52.18%)
 * Total Accepted:    9.4K
 * Total Submissions: 18K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * 给定一个二进制数组， 计算其中最大连续1的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,1,0,1,1,1]
 * 输出: 3
 * 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 *
 *
 * 注意：
 *
 *
 * 输入的数组只包含 0 和1。
 * 输入数组的长度是正整数，且不超过 10,000。
 *
 *
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaxConsecutiveOnes = function(nums) {
  let res = 0
  let cur = 0
  for (const n of nums) {
    if (n & 1) {
      cur++
      res = Math.max(res, cur)
    } else {
      cur = 0
    }
  }
  return res
}
