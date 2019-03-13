/*
 * @lc app=leetcode.cn id=5 lang=javascript
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (24.64%)
 * Total Accepted:    40.1K
 * Total Submissions: 162.6K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  const arr = ['$', '#']
  for (let i = 0; i < s.length; i++) {
    arr.push(s[i], '#')
  }
  arr.push('%')

  let res = ''
  let maxPoint = 0
  let rightMax = 0
  let middle = 0
  const lens = arr.slice().fill(1)

  for (let i = 0; i < arr.length; i++) {
    if (rightMax > i) {
      lens[i] = Math.min(rightMax - i, lens[2 * middle - i])
    }

    while (arr[i - lens[i]] === arr[i + lens[i]]) {
      lens[i]++
    }

    if (lens[i] + i > rightMax) {
      middle = i
      rightMax = lens[i] + i
    }

    if (lens[i] > lens[maxPoint]) {
      maxPoint = i
    }
  }

  for (
    let i = maxPoint - (lens[maxPoint] - 2);
    i < maxPoint + lens[maxPoint] - 1;
    i += 2
  ) {
    res += arr[i]
  }
  return res
}
