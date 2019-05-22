/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (59.33%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    12.4K
 * Total Submissions: 21K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */
func minimumTotal(triangle [][]int) int {
	res := 1 << 20
	for i := range triangle {
		for j := range triangle[i] {
			if i == 0 {
				continue
			}
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j-1]+triangle[i][j],
					triangle[i-1][j]+triangle[i][j])
			}
		}
	}
	for _, n := range triangle[len(triangle)-1] {
		res = min(res, n)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
