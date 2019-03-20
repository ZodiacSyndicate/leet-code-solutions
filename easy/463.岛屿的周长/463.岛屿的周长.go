/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 *
 * https://leetcode-cn.com/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (59.24%)
 * Total Accepted:    4.7K
 * Total Submissions: 7.9K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * 给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
 *
 * 网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
 *
 * 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100
 * 。计算这个岛屿的周长。
 *
 *
 *
 * 示例 :
 *
 * 输入:
 * [[0,1,0,0],
 * ⁠[1,1,1,0],
 * ⁠[0,1,0,0],
 * ⁠[1,1,0,0]]
 *
 * 输出: 16
 *
 * 解释: 它的周长是下面图片中的 16 个黄色的边：
 *
 *
 *
 *
 */
func islandPerimeter(grid [][]int) int {
	ans := 0
	for y, _ := range grid {
		for x, v := range grid[y] {
			n := 0
			if v == 1 {
				n = 4
				if y > 0 && grid[y-1][x] == 1 {
					n--
				}
				if y < len(grid)-1 && grid[y+1][x] == 1 {
					n--
				}
				if x > 0 && grid[y][x-1] == 1 {
					n--
				}
				if x < len(grid[y])-1 && grid[y][x+1] == 1 {
					n--
				}
			}
			ans += n
		}
	}
	return ans
}

