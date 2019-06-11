/*
 * @lc app=leetcode.cn id=131 lang=javascript
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (62.32%)
 * Likes:    115
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 12.3K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */
/**
 * @param {string} s
 * @return {string[][]}
 */
var partition = function(s) {
  const res = []
  function backtrack(str, arr) {
    if (isPalindrome(str)) {
      const temp = [...arr]
      temp.push(str)
      res.push(temp)
    }
    for (let i = 1; i < str.length; i++) {
      const temp = [...arr]
      const s = str.slice(0, i)
      if (!isPalindrome(s)) continue
      temp.push(s)
      backtrack(str.slice(i), temp)
    }
  }
  backtrack(s, [])
  return res
}

function isPalindrome(s) {
  if (s.length === 1) return true
  let i = 0
  let j = s.length - 1
  while (i <= j) {
    if (s[i] !== s[j]) return false
    i++
    j--
  }
  return true
}
