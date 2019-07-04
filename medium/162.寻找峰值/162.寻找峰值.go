/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 *
 * https://leetcode-cn.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (39.89%)
 * Likes:    59
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 28.8K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 峰值元素是指其值大于左右相邻值的元素。
 *
 * 给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
 *
 * 数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
 *
 * 你可以假设 nums[-1] = nums[n] = -∞。
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1]
 * 输出: 2
 * 解释: 3 是峰值元素，你的函数应该返回其索引 2。
 *
 * 示例 2:
 *
 * 输入: nums = [1,2,1,3,5,6,4]
 * 输出: 1 或 5
 * 解释: 你的函数可以返回索引 1，其峰值元素为 2；
 * 或者返回索引 5， 其峰值元素为 6。
 *
 *
 * 说明:
 *
 * 你的解法应该是 O(logN) 时间复杂度的。
 *
 */
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	if l == r {
		return l
	}
	if nums[l] > nums[l+1] {
		return l
	}
	if nums[r] > nums[r-1] {
		return r
	}
	for l <= r {
		mid := l + (r-l)/2
		if mid != 0 && mid != len(nums)-1 && nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		} else if mid > 0 && nums[mid] < nums[mid-1] {
			r = mid - 1
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			l = mid + 1
		}
	}
	return 0
}

