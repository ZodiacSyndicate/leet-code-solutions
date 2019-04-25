/*
 * @lc app=leetcode.cn id=60 lang=javascript
 *
 * [60] 第k个排列
 */
/**
 * @param {number} n
 * @param {number} k
 * @return {string}
 */
var getPermutation = function(n, k) {
  const arr = Array.from({ length: n }, (_, i) => i + 1)
  const res = []
  let pos = k - 1
  let fac = arr.reduce((a, b) => a * b)
  for (let j = n; j >= 1; --j) {
    fac /= j
    res.push(arr.splice(~~(pos / fac), 1)[0])
    pos %= fac
  }
  return res.join('')
}
