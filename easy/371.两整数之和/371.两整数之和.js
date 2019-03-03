/*
 * @lc app=leetcode.cn id=371 lang=javascript
 *
 * [371] 两整数之和
 *
 * https://leetcode-cn.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (57.99%)
 * Total Accepted:    7.4K
 * Total Submissions: 12.9K
 * Testcase Example:  '1\n2'
 *
 * 不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
 *
 * 示例 1:
 *
 * 输入: a = 1, b = 2
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: a = -2, b = 3
 * 输出: 1
 *
 */
/**
 * @param {number} a
 * @param {number} b
 * @return {number}
 */
var getSum = function(a, b) {
  const sum = a ^ b
  const carry = (a & b) << 1
  if (carry != 0) return getSum(sum, carry)
  return sum
}
