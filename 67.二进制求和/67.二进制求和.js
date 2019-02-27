/*
 * @lc app=leetcode.cn id=67 lang=javascript
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (46.14%)
 * Total Accepted:    15.9K
 * Total Submissions: 34.3K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
  let res = ''
  let flag = false
  while (a || b) {
    const [a1, b1] = [Number(a.slice(-1)), Number(b.slice(-1))]
    if (flag) {
      if (a1 + b1 + 1 < 2) {
        flag = false
      }
      res = `${(a1 + b1 + 1) % 2}${res}`
    } else {
      if (a1 + b1 === 2) {
        flag = true
      }
      res = `${(a1 + b1) % 2}${res}`
    }

    ;[a, b] = [a.slice(0, a.length - 1), b.slice(0, b.length - 1)]
  }
  if (flag) res = `1${res}`
  return res
}
