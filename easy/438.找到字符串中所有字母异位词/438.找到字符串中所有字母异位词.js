/*
 * @lc app=leetcode.cn id=438 lang=javascript
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (36.21%)
 * Total Accepted:    3.5K
 * Total Submissions: 9.7K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 *
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 *
 * 说明：
 *
 *
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * s: "cbaebabacd" p: "abc"
 *
 * 输出:
 * [0, 6]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s: "abab" p: "ab"
 *
 * 输出:
 * [0, 1, 2]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 *
 *
 */
/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */
var findAnagrams = function(s, p) {
  const map = []
  let map1 = []
  const res = []
  const len = p.length
  for (let i = 0; i < len; i++) {
    const index = p[i].charCodeAt()
    if (map[index]) {
      map[index]++
      map1[index]++
    } else {
      map[index] = 1
      map1[index] = 1
    }
  }
  for (let i = 0; i + len <= s.length; i++) {
    const str = s.slice(i, i + len)
    let flag = true
    for (let j = 0; j < len; j++) {
      const index = str[j].charCodeAt()
      if (map1[index]) {
        map1[index]--
        if (map1[index] < 0) {
          flag = false
          break
        }
      } else {
        flag = false
        break
      }
    }
    if (flag) {
      res.push(i)
    }
    map1 = [...map]
  }
  return res
}
