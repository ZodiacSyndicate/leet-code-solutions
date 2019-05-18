/*
 * @lc app=leetcode.cn id=89 lang=javascript
 *
 * [89] 格雷编码
 */
/**
 * @param {number} n
 * @return {number[]}
 */
var grayCode = function(n) {
  if (n === 0) return [0]
  const res = grayCode(n - 1)
  const mask = 1 << (n - 1)
  for (let i = res.length - 1; i >= 0; i--) {
    res.push(res[i] | mask)
  }
  return res
}
