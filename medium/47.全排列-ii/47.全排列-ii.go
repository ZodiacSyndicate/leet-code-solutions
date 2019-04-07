/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (51.19%)
 * Total Accepted:    9.4K
 * Total Submissions: 18.4K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 */
import "sort"
func permuteUnique(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	ans := [][]int{}
	var perm func(nums []int, start int, end int)
	perm = func(nums []int, start int, end int) {
		if start == end {
			ans = append(ans, nums)
		} else {
			for i := start; i <= end; i++ {
				if i != start && nums[start] == nums[i] {
					continue
				}
				nums[start], nums[i] = nums[i], nums[start]
				arr := make([]int, len(nums))
				copy(arr, nums)
				perm(arr, start + 1, end)
			}
		}
	}
	perm(nums, 0, len(nums)-1)
	return ans
}

