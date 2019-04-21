/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (34.56%)
 * Total Accepted:    12.2K
 * Total Submissions: 35.2K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 */
func spiralOrder(matrix [][]int) []int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := []int{}
	x, y, i, flag := 0, 0, 0, false
	for {
		if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[x]) && matrix[x][y] != -100000 {
			flag = false
			res = append(res, matrix[x][y])
			matrix[x][y] = -100000
		} else {
			if flag {
				break
			}
			flag = true
			x -= directions[i][0]
			y -= directions[i][1]
			i = (i + 1) % 4
		}
		x += directions[i][0]
		y += directions[i][1]
	}
	return res
}

