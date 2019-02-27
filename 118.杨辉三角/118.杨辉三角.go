/*
 * @lc app=leetcode.cn id=118 lang=golang
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
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := [][]int{[]int{1}}
	arr := []int{1}
	for numRows > 1 {
		cur := []int{}
		for i := 0; i <= len(arr); i++ {
			if i == 0 || i == len(arr) {
				cur = append(cur, 1)
			} else {
				cur = append(cur, arr[i]+arr[i-1])
			}
		}
		res = append(res, cur)
		arr = cur
		numRows--
	}
	return res
}
