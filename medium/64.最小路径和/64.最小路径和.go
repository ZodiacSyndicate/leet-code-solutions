/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
func minPathSum(grid [][]int) int {
	x, y := len(grid), len(grid[0])
	dp := make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, y)
		for j := 0; j < y; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[x-1][y-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

