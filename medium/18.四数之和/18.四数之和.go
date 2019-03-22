/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (34.65%)
 * Total Accepted:    14K
 * Total Submissions: 40.3K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	return sum(nums, target, 4)
}

func sum(nums []int, target int, flag int) [][]int {
	var res [][]int
	if flag == 1 {
		if indexOf(nums, target) >= 0 {
			return [][]int{{target}}
		}
	} else {
		for i := range nums {
			if i != 0 && nums[i] == nums[i-1] {
				continue
			}
			temp := sum(nums[i+1:], target-nums[i], flag-1)
			if len(temp) != 0 {
				for _, arr := range temp {
					arr = append(arr, nums[i])
					res = append(res, arr)
				}
			}
		}
	}
	return res
}

func indexOf(arr []int, num int) int {
	for i, n := range arr {
		if n == num {
			return i
		}
	}
	return -1
}

