/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (53.57%)
 * Total Accepted:    9.7K
 * Total Submissions: 18K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */
func getRow(rowIndex int) []int {
	res := []int{1}
	for rowIndex > 0 {
		cur := []int{}
		for i := 0; i <= len(res); i++ {
			if i == 0 || i == len(res) {
				cur = append(cur, 1)
			} else {
				cur = append(cur, res[i]+res[i-1])
			}
		}
		res = cur
		rowIndex--
	}
	return res
}
