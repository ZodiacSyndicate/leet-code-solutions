/*
 * @lc app=leetcode.cn id=345 lang=javascript
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (45.92%)
 * Total Accepted:    6.6K
 * Total Submissions: 14.4K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 * 示例 1:
 *
 * 输入: "hello"
 * 输出: "holle"
 *
 *
 * 示例 2:
 *
 * 输入: "leetcode"
 * 输出: "leotcede"
 *
 * 说明:
 * 元音字母不包含字母"y"。
 *
 */
/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function(s) {
  let l = 0
  let r = s.length - 1
  const arr = s.split('')
  while (l <= r) {
    while (!isVowel(arr[l]) && l <= r) {
      l++
    }
    while (!isVowel(arr[r]) && l <= r) {
      r--
    }
    if (l <= r) {
      ;[arr[l], arr[r]] = [arr[r], arr[l]]
    }
    l++
    r--
  }
  return arr.join('')
}

function isVowel(s) {
  return (
    s === 'a' ||
    s === 'e' ||
    s === 'i' ||
    s === 'o' ||
    s === 'u' ||
    s === 'A' ||
    s === 'E' ||
    s === 'I' ||
    s === 'O' ||
    s === 'U'
  )
}
