/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (67.44%)
 * Total Accepted:    19.9K
 * Total Submissions: 29.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */
func permute(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(nums []int, start int, end int)
	backtrack = func(nums []int, start int, end int) {
		if start == end {
			result = append(result, nums)
		} else {
			for i := start; i < end; i++ {
				nums[i], nums[start] = nums[start], nums[i]
				arr := make([]int, len(nums))
				copy(arr, nums)
				backtrack(arr, start + 1, end)
				nums[i], nums[start] = nums[start], nums[i]
			}
		}
	}
	backtrack(nums, 0, len(nums) )
	return result
}


