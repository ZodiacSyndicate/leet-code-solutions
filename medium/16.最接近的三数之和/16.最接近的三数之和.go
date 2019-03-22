/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (39.31%)
 * Total Accepted:    18.9K
 * Total Submissions: 48.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */
import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	closest := nums[0] + nums[1] + nums[2]
	for i, n := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if abs(threeSum-target) < abs(closest-target) {
				closest = threeSum
			}
			if threeSum < target {
				l++
			} else if threeSum > target {
				r--
			} else {
				return threeSum
			}
		}
	}
	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

