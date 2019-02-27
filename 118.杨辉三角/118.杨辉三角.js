/*
 * @lc app=leetcode.cn id=118 lang=javascript
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (59.45%)
 * Total Accepted:    15.9K
 * Total Submissions: 26.6K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */
/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
  if (!numRows) return []
  let n = 1
  const result = [[1]]
  while (n < numRows) {
    const cur = result[result.length - 1]
    const arr = []
    for (let i = 0; i <= cur.length; i++) {
      arr[i] = (cur[i] ? cur[i] : 0) + (cur[i - 1] ? cur[i - 1] : 0)
    }
    result.push(arr)
    n++
  }
  return result
}
