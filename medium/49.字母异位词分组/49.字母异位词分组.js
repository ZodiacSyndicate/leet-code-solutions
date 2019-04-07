/*
 * @lc app=leetcode.cn id=49 lang=javascript
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (54.77%)
 * Total Accepted:    13.2K
 * Total Submissions: 24.1K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */
/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
  if (strs.length === 0) return []
  const ans = {}
  const count = Array.from({ length: 26 })
  for (let s of strs) {
    count.fill(0)
    for (let c of s.split('')) count[c.charCodeAt() - 'a'.charCodeAt()]++

    let key = ''
    for (let i = 0; i < 26; i++) {
      key += '#'
      key += count[i]
    }
    if (key in ans) {
      ans[key].push(s)
    } else {
      ans[key] = [s]
    }
  }
  return Object.values(ans)
}
