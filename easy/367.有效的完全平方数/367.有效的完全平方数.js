/*
 * @lc app=leetcode.cn id=367 lang=javascript
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (39.57%)
 * Total Accepted:    6.3K
 * Total Submissions: 15.8K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */
/**
 * @param {number} num
 * @return {boolean}
 */
var isPerfectSquare = function(num) {
  let l = 0
  let r = num
  while (l <= r) {
    const mid = (l + (r - l) / 2) | 0
    if (mid ** 2 === num) {
      return true
    } else if (mid ** 2 < num) {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return false
}
