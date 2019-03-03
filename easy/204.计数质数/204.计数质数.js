/*
 * @lc app=leetcode.cn id=204 lang=javascript
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (25.80%)
 * Total Accepted:    12.4K
 * Total Submissions: 47.4K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */
/**
 * @param {number} n
 * @return {number}
 */
var countPrimes = function(n) {
  const isPrime = []
  for (let i = 2; i < n; i++) {
    isPrime[i] = true
  }
  for (let i = 2; i * i < n; i++) {
    if (!isPrime[i]) continue
    for (let j = i * i; j < n; j += i) {
      isPrime[j] = false
    }
  }
  let count = 0
  for (let i = 2; i < n; i++) {
    if (isPrime[i]) count++
  }
  return count
}
