/*
 * @lc app=leetcode.cn id=459 lang=javascript
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (38.79%)
 * Total Accepted:    4K
 * Total Submissions: 10.2K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */
/**
 * @param {string} s
 * @return {boolean}
 */
var repeatedSubstringPattern = function(s) {
  const arr = []
  let prev = 0
  for (let i = 1; i < s.length; i++) {
    if (s[i] === s[0]) {
      arr.push(i)
    }
    if (s.length % (i - prev) !== 0) {
      return false
    }
    prev = i
  }
  if (arr.length === 0) return false
  for (let item of arr) {
    if (s.slice(0, item).repeat(s.length / item) === s) return true
  }
  return false
}
