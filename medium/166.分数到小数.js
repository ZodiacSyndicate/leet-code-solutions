/*
 * @lc app=leetcode.cn id=166 lang=javascript
 *
 * [166] 分数到小数
 *
 * https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (24.01%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    3.2K
 * Total Submissions: 13.4K
 * Testcase Example:  '1\n2'
 *
 * 给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
 *
 * 如果小数部分为循环小数，则将循环的部分括在括号内。
 *
 * 示例 1:
 *
 * 输入: numerator = 1, denominator = 2
 * 输出: "0.5"
 *
 *
 * 示例 2:
 *
 * 输入: numerator = 2, denominator = 1
 * 输出: "2"
 *
 * 示例 3:
 *
 * 输入: numerator = 2, denominator = 3
 * 输出: "0.(6)"
 *
 *
 */
/**
 * @param {number} numerator
 * @param {number} denominator
 * @return {string}
 */
var fractionToDecimal = function(numerator, denominator) {
  if (numerator === 0) return '0'
  let fraction = ''
  if ((numerator < 0) ^ (denominator < 0)) {
    fraction += '-'
  }
  const dividend = Math.abs(numerator)
  const divisor = Math.abs(denominator)
  fraction += Math.floor(dividend / divisor).toString()
  let remainder = dividend % divisor
  if (remainder === 0) return fraction
  fraction += '.'
  const map = {}
  while (remainder != 0) {
    if (remainder in map) {
      fraction =
        fraction.slice(0, map[remainder]) +
        '(' +
        fraction.slice(map[remainder]) +
        ')'
      break
    }
    map[remainder] = fraction.length
    remainder *= 10
    fraction += Math.floor(remainder / divisor).toString()
    remainder %= divisor
  }
  return fraction
}
