/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	x, y := len(matrix), len(matrix[0])
	l, r := 0, x*y-1
	for l <= r {
		mid := (l + r) / 2
		ix, iy := mid/y, mid%y
		if matrix[ix][iy] == target {
			return true
		} else if matrix[ix][iy] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

