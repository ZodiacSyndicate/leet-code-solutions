/*
 * @lc app=leetcode.cn id=168 lang=javascript
 *
 * [168] Excel表列名称
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (29.78%)
 * Total Accepted:    5.1K
 * Total Submissions: 17K
 * Testcase Example:  '1'
 *
 * 给定一个正整数，返回它在 Excel 表中相对应的列名称。
 *
 * 例如，
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "A"
 *
 *
 * 示例 2:
 *
 * 输入: 28
 * 输出: "AB"
 *
 *
 * 示例 3:
 *
 * 输入: 701
 * 输出: "ZY"
 *
 *
 */
/**
 * @param {number} n
 * @return {string}
 */
var convertToTitle = function(n) {
  const arr = ['Z']
  let res = ''
  for (let i = 65; i < 90; i++) {
    arr.push(String.fromCharCode(i))
  }
  while (n !== 0) {
    const num = n % 26
    res = `${arr[num]}${res}`
    n = (n - num) / 26
    if (num === 0) {
      n--
    }
  }
  return res
}
