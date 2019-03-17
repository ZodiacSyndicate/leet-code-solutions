/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (49.41%)
 * Total Accepted:    2.7K
 * Total Submissions: 5.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。
 *
 * 示例:
 *
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 3
 *
 * 解释:
 * 只需要3次移动（注意每次移动会增加两个元素的值）：
 *
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 *
 *
 */
import "math"

func minMoves(nums []int) int {
	sum := 0
	min := float64(^uint(0))
	for _, n := range nums {
		sum += n
		min = math.Min(min, float64(n))
	}
	return sum - int(min)*len(nums)
}

