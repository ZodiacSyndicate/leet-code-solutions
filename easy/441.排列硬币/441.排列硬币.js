/*
 * @lc app=leetcode.cn id=441 lang=javascript
 *
 * [441] 排列硬币
 *
 * https://leetcode-cn.com/problems/arranging-coins/description/
 *
 * algorithms
 * Easy (35.79%)
 * Total Accepted:    5.3K
 * Total Submissions: 14.8K
 * Testcase Example:  '5'
 *
 * 你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
 *
 * 给定一个数字 n，找出可形成完整阶梯行的总行数。
 *
 * n 是一个非负整数，并且在32位有符号整型的范围内。
 *
 * 示例 1:
 *
 *
 * n = 5
 *
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤
 *
 * 因为第三行不完整，所以返回2.
 *
 *
 * 示例 2:
 *
 *
 * n = 8
 *
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 *
 * 因为第四行不完整，所以返回3.
 *
 *
 */
/**
 * @param {number} n
 * @return {number}
 */
var arrangeCoins = function(n) {
  if (!n) return 0
  let l = 1
  let r = ~~(n / 2) + 1
  while (l <= r) {
    const mid = ~~((l + r) / 2)
    if (sum(mid) <= n && sum(mid + 1) > n) {
      return mid
    } else if (sum(mid) < n) {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
}

function sum(n) {
  return ((1 + n) * n) / 2
}
