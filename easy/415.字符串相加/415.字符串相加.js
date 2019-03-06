/*
 * @lc app=leetcode.cn id=415 lang=javascript
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (42.83%)
 * Total Accepted:    4.9K
 * Total Submissions: 11.3K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */
/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function(num1, num2) {
  let l1 = num1.length - 1
  let l2 = num2.length - 1
  let carry = 0
  let ans = ''
  while (l1 >= 0 || l2 >= 0) {
    const n1 = l1 >= 0 ? Number(num1[l1]) : 0
    const n2 = l2 >= 0 ? Number(num2[l2]) : 0
    let val = carry + n1 + n2
    carry = ~~(val / 10)
    val %= 10
    ans = `${val}${ans}`
    l1--
    l2--
  }
  return carry ? `${carry}${ans}` : ans
}
