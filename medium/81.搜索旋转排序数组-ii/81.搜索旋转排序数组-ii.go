/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (32.76%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 18K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 * 
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 * 
 * 示例 1:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 * 
 * 进阶:
 * 
 * 
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 * 
 * 
 */
func search(nums []int, target int) bool {
  return mySearch(nums, target, 0, len(nums) - 1)
}

func mySearch(nums []int, target, i, j int) bool {
	if j < 0 {
		return false
	}
	mid := (i + j) / 2
	if nums[mid] == target || nums[i] == target || nums[j] == target {
		return true
	} else if mid == i {
		return false
	} else if nums[i] == nums[mid] && nums[mid] == nums[j] {
		return mySearch(nums, target, i, mid - 1) || mySearch(nums, target, mid, j)
	} else if nums[mid] >= nums[i] {
		if target > nums[i] && target < nums[mid] {
			return mySearch(nums, target, i, mid - 1)
		}
		return mySearch(nums, target, mid + 1, j)
	} else {
		if target > nums[mid] && target < nums[i] {
			return mySearch(nums, target, mid + 1, j)
		}
		return mySearch(nums, target, i, mid - 1)
	}
}

