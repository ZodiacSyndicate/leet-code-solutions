/*
 * @lc app=leetcode.cn id=93 lang=javascript
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (43.43%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    8K
 * Total Submissions: 18.1K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 *
 * 示例:
 *
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 *
 */
/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function(s) {
  const res = []
  function backtrack(str, temp = [], k = 0) {
    if (temp.length > 4) return
    if (k === 3) {
      if (Number(str) <= 255 && str.length <= 3) {
        temp.push(str)
        res.push(temp.join('.'))
      }
      return
    }
    for (let i = 1; i <= 3; i++) {
      if (Number(str.slice(0, i)) > 255) return
      temp.push(str.slice(0, i))
      backtrack(str.slice(i), [...temp], k + 1)
    }
  }
  backtrack(s, [], 0)
  return res
}
