/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (21.45%)
 * Total Accepted:    41.5K
 * Total Submissions: 193.6K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
import "sort"

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	c := len(nums) - 1
	var ans [][]int
	for c >= 2 {
		a, b := 0, c-1
		for a < b {
			sum := nums[a] + nums[b]
			if sum < -nums[c] {
				a++
			} else if sum > -nums[c] {
				b--
			} else {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
				a++
				for a < b && nums[a-1] == nums[a] {
					a++
				}
				b--
				for a < b && nums[b+1] == nums[b] {
					b--
				}
			}
		}
		c--
		for c >= 2 && nums[c+1] == nums[c] {
			c--
		}
	}
	return ans
}

