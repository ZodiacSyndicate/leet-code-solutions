/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (34.18%)
 * Likes:    200
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 33.4K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */
func maxProduct(nums []int) int {
	res := nums[0]
	min := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		max = maxi(nums[i], max*nums[i])
		min = mini(nums[i], min*nums[i])
		res = maxi(res, max)
	}
	return res
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}
