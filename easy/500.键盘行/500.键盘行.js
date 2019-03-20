/*
 * @lc app=leetcode.cn id=500 lang=javascript
 *
 * [500] 键盘行
 *
 * https://leetcode-cn.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (64.94%)
 * Total Accepted:    5.7K
 * Total Submissions: 8.8K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * 给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
 *
 *
 *
 *
 *
 *
 *
 * 示例：
 *
 * 输入: ["Hello", "Alaska", "Dad", "Peace"]
 * 输出: ["Alaska", "Dad"]
 *
 *
 *
 *
 * 注意：
 *
 *
 * 你可以重复使用键盘上同一字符。
 * 你可以假设输入的字符串将只包含字母。
 *
 */
/**
 * @param {string[]} words
 * @return {string[]}
 */

const line1 = 'qwertyuiopQWERTYUIOP'.split('').reduce((map, s) => {
  map[s] = 1
  return map
}, {})
const line2 = 'asdfghjklASDFGHJKL'.split('').reduce((map, s) => {
  map[s] = 1
  return map
}, {})
const line3 = 'zxcvbnmZXCVBNM'.split('').reduce((map, s) => {
  map[s] = 1
  return map
}, {})
var findWords = function(words) {
  const ans = []
  for (let word of words) {
    let line
    const s = word[0]
    if (s in line1) {
      line = line1
    } else if (s in line2) {
      line = line2
    } else {
      line = line3
    }
    let flag = true
    for (let i = 0; i < word.length; i++) {
      if (!(word[i] in line)) {
        flag = false
        break
      }
    }
    if (flag) {
      ans.push(word)
    }
  }
  return ans
}
