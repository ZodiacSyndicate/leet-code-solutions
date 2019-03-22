/*
 * @lc app=leetcode.cn id=17 lang=javascript
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (47.98%)
 * Total Accepted:    18.1K
 * Total Submissions: 37.6K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */
/**
 * @param {string} digits
 * @return {string[]}
 */
const list = [
  null,
  null,
  'abc',
  'def',
  'ghi',
  'jkl',
  'mno',
  'pqrs',
  'tuv',
  'wxyz'
]
var letterCombinations = function(digits) {
  let res = []
  if (!digits) return res
  dfs(res, '', digits, list, 0)
  return res
}

function dfs(res, str, digits, list, i) {
  if (str.length === digits.length) {
    res.push(str)
    return
  }
  const temp = list[digits[i]]
  for (let j = 0; j < temp.length; j++) {
    str += temp[j]
    dfs(res, str, digits, list, i + 1)
    str = str.slice(0, str.length - 1)
  }
  return
}
