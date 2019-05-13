/*
 * @lc app=leetcode.cn id=77 lang=javascript
 *
 * [77] 组合
 */
/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
var combine = function(n, k) {
  if (k > n || k === 0) return []
  if (k === 1) return Array.from({ length: n }, (_, i) => [i + 1])
  if (k === n) return [Array.from({ length: n }, (_, i) => i + 1)]
  const res = combine(n - 1, k)
  for (const item of combine(n - 1, k - 1)) {
    item.push(n)
    res.push(item)
  }
  return res
}
