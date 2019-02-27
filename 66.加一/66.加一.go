/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (37.33%)
 * Total Accepted:    35.9K
 * Total Submissions: 96.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */
func plusOne(digits []int) []int {
	flag := true
	for n := len(digits) - 1; n >= 0; n-- {
		if flag {
			if digits[n]+1 < 10 {
				flag = false
			}
			digits[n] = (digits[n] + 1) % 10
		} else {
			return digits
		}
	}
	if flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}
