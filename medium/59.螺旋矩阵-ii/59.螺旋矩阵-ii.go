/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (71.53%)
 * Total Accepted:    8.5K
 * Total Submissions: 11.8K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */
func generateMatrix(n int) [][]int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		arr := []int{}
		for j := 0; j < n; j++ {
			arr = append(arr, -1002345)
		}
		matrix = append(matrix, arr)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, i, k, flag := 0, 0, 0, 1, false
	for {
		if x < n && y < n && x >= 0 && y >= 0 && matrix[x][y] == -1002345 {
			flag = false
			matrix[x][y] = k
			k++
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
	return matrix
}

