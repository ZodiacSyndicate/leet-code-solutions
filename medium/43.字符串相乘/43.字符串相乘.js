/*
 * @lc app=leetcode.cn id=43 lang=javascript
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (37.97%)
 * Total Accepted:    12.2K
 * Total Submissions: 32.3K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 示例 1:
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * 说明：
 *
 *
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 *
 *
 */

/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var multiply = function(num1, num2) {
  const n1 = num1.length - 1
  const n2 = num2.length - 1
  if (n1 < 0 || n2 < 0) return ''
  const mul = Array.from({ length: n1 + n2 + 2 }).fill(0)
  for (let i = n1; i >= 0; i--) {
    for (let j = n2; j >= 0; j--) {
      let bitMul =
        (num1.charCodeAt(i) - '0'.charCodeAt()) *
        (num2.charCodeAt(j) - '0'.charCodeAt())
      bitMul += mul[i + j + 1]
      mul[i + j] += ~~(bitMul / 10)
      mul[i + j + 1] = bitMul % 10
    }
  }

  let i = 0
  while (i < mul.length - 1 && mul[i] === 0) i++
  return mul.slice(i).join('')
}
