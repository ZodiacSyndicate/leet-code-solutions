/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 *
 * https://leetcode-cn.com/problems/number-of-boomerangs/description/
 *
 * algorithms
 * Easy (52.92%)
 * Total Accepted:    3.2K
 * Total Submissions: 5.9K
 * Testcase Example:  '[[0,0],[1,0],[2,0]]'
 *
 * 给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k
 * 之间的距离相等（需要考虑元组的顺序）。
 *
 * 找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
 *
 * 示例:
 *
 *
 * 输入:
 * [[0,0],[1,0],[2,0]]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
 *
 *
 */
func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i, p1 := range points {
		m := map[int]int{}
		for j, p2 := range points {
			if i != j {
				a1, b1, a2, b2 := p1[0], p1[1], p2[0], p2[1]
				dis := (a1-a2)*(a1-a2) + (b1-b2)*(b1-b2)
				if _, ok := m[dis]; ok {
					m[dis]++
				} else {
					m[dis] = 1
				}
			}
		}
		for _, val := range m {
			if val > 1 {
				ans += val * (val - 1)
			}
		}
	}
	return ans
}

