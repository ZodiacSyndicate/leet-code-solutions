/*
 * @lc app=leetcode.cn id=434 lang=javascript
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (28.93%)
 * Total Accepted:    3.9K
 * Total Submissions: 13.4K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 *
 *
 */
/**
 * @param {string} s
 * @return {number}
 */
var countSegments = function(s) {
  if (!s.trim()) return 0
  return s.split(' ').filter(item => item).length
}
