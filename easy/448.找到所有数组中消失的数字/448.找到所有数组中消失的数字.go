/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (46.40%)
 * Total Accepted:    5.6K
 * Total Submissions: 12.1K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 *
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 *
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 *
 * 示例:
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [5,6]
 *
 *
 */

func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	for _, n := range nums {
		if n >= 0 {
			if nums[n-1] > 0 {
				nums[n-1] = -nums[n-1]
			}
		} else {
			if nums[-n-1] > 0 {
				nums[-n-1] = -nums[-n-1]
			}
		}
	}
	for i, n := range nums {
		if n > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

