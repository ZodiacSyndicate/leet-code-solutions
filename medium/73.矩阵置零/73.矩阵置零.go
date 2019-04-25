/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
func setZeroes(matrix [][]int) {
	rowFlag := false
	for _, n := range matrix[0] {
		if n == 0 {
			rowFlag = true
			break
		}
	}

	colFlag := false
	for i := range matrix {
		if matrix[i][0] == 0 {
			colFlag = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := range matrix {
				matrix[j][i] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := range matrix[0] {
				matrix[i][j] = 0
			}
		}
	}

	if rowFlag {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}

	if colFlag {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}

