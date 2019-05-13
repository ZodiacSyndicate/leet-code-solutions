/*
 * @lc app=leetcode.cn id=78 lang=javascript
 *
 * [78] 子集
 */
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
  const result = [[]]
  nums.forEach(n => {
    result.forEach(item => {
      result.push(item.concat(n))
    })
  })
  return result
}
