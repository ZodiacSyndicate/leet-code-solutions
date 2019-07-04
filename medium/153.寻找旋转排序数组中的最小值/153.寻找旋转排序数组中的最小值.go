/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (49.28%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    10.6K
 * Total Submissions: 21.5K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */
func findMin(nums []int) int {
	return find(nums, 0, len(nums)-1)
}

func find(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	mid := l + (r-l)/2
	if nums[l] < nums[r] {
		return nums[l]
	}
	if mid == 0 && nums[mid] < nums[mid+1] {
		return nums[mid]
	}
	if mid == len(nums)-1 && nums[mid] < nums[mid-1] {
		return nums[mid]
	}
	if mid != 0 && mid != len(nums)-1 && nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
		return nums[mid]
	}
	if nums[mid] > nums[r] {
		return find(nums, mid+1, r)
	}
	if nums[mid] < nums[l] {
		return find(nums, l, mid-1)
	}
	return 0
}

