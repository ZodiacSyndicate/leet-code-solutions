/*
 * @lc app=leetcode.cn id=263 lang=javascript
 *
 * [263] 丑数
 *
 * https://leetcode-cn.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (44.39%)
 * Total Accepted:    6.3K
 * Total Submissions: 14.2K
 * Testcase Example:  '6'
 *
 * 编写一个程序判断给定的数是否为丑数。
 *
 * 丑数就是只包含质因数 2, 3, 5 的正整数。
 *
 * 示例 1:
 *
 * 输入: 6
 * 输出: true
 * 解释: 6 = 2 × 3
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: true
 * 解释: 8 = 2 × 2 × 2
 *
 *
 * 示例 3:
 *
 * 输入: 14
 * 输出: false
 * 解释: 14 不是丑数，因为它包含了另外一个质因数 7。
 *
 * 说明：
 *
 *
 * 1 是丑数。
 * 输入不会超过 32 位有符号整数的范围: [−231,  231 − 1]。
 *
 *
 */
/**
 * @param {number} num
 * @return {boolean}
 */
var isUgly = function(num) {
  if (!num) return false
  while (num % 2 === 0) {
    num /= 2
  }
  while (num % 3 === 0) {
    num /= 3
  }
  while (num % 5 === 0) {
    num /= 5
  }
  return num === 1
}
