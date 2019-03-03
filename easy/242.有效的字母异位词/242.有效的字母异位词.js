/*
 * @lc app=leetcode.cn id=242 lang=javascript
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (49.88%)
 * Total Accepted:    20.3K
 * Total Submissions: 40.4K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
  let i = 0
  const x = Array.from({ length: 26 }).fill(0)
  const y = Array.from({ length: 26 }).fill(0)
  for (i = 0; i < s.length; i++) {
    x[s[i].charCodeAt() - 'a'.charCodeAt()]++
  }
  for (i = 0; i < t.length; i++) {
    y[t[i].charCodeAt() - 'a'.charCodeAt()]++
  }
  for (i = 0; i < 26; i++) {
    if (x[i] !== y[i]) return false
  }
  return true
}
