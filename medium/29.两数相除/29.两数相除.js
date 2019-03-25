/*
 * @lc app=leetcode.cn id=29 lang=javascript
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (17.61%)
 * Total Accepted:    9.4K
 * Total Submissions: 53.3K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 *
 * 示例 1:
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 *
 * 说明:
 *
 *
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 *
 *
 */
/**
 * @param {number} dividend
 * @param {number} divisor
 * @return {number}
 */
const INT_MAX = 2 ** 31 - 1
const INT_MIN = -(2 ** 31)
var divide = function(dividend, divisor) {
  let ans = 0
  let [up, down] = [Math.abs(dividend), Math.abs(divisor)]
  while (up >= down) {
    let count = 1
    let base = down
    while (up > base << 1 && base << 1 > 0) {
      count <<= 1
      base <<= 1
    }
    ans += count
    up -= base
  }
  ans = (dividend < 0) ^ (divisor < 0) ? -ans : ans
  return ans > INT_MAX || ans < INT_MIN ? INT_MAX : ans
}
