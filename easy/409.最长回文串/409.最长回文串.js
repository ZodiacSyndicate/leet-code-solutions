/*
 * @lc app=leetcode.cn id=409 lang=javascript
 *
 * [409] 最长回文串
 *
 * https://leetcode-cn.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (45.63%)
 * Total Accepted:    3.6K
 * Total Submissions: 7.8K
 * Testcase Example:  '"abccccdd"'
 *
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *
 * 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 *
 * 注意:
 * 假设字符串的长度不会超过 1010。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "abccccdd"
 *
 * 输出:
 * 7
 *
 * 解释:
 * 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 *
 */
/**
 * @param {string} s
 * @return {number}
 */
var longestPalindrome = function(s) {
  const map = {}
  let len = 0
  let flag = false
  for (let i = 0; i < s.length; i++) {
    const str = s[i]
    if (map[str]) {
      map[str]++
    } else {
      map[str] = 1
    }
  }
  for (const key of Object.keys(map)) {
    if (map[key] % 2 === 0) {
      len += map[key]
    } else {
      len += map[key] - 1
      flag = true
    }
  }
  return flag ? len + 1 : len
}
