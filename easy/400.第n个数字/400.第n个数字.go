/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第N个数字
 *
 * https://leetcode-cn.com/problems/nth-digit/description/
 *
 * algorithms
 * Easy (30.10%)
 * Total Accepted:    2.1K
 * Total Submissions: 7K
 * Testcase Example:  '3'
 *
 * 在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。
 *
 * 注意:
 * n 是正数且在32为整形范围内 ( n < 231)。
 *
 * 示例 1:
 *
 *
 * 输入:
 * 3
 *
 * 输出:
 * 3
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * 11
 *
 * 输出:
 * 0
 *
 * 说明:
 * 第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
 *
 *
 */

func findNthDigit(n int) int {
	cnt := 1
	sum := 9
	tmp := 9
	isOverFlow := 0
	for n > sum {
		tmp *= 10
		cnt++
		if sum+tmp*cnt < sum {
			isOverFlow = 1
			break
		}
		sum += tmp * cnt
	}
	if isOverFlow == 0 {
		sum -= tmp * cnt
	}
	num := tmp/9 + (n-sum-1)/cnt
	idx := cnt - ((n-sum-1)%cnt + 1)
	res := -1
	for idx >= 0 {
		idx -= 1
		res = num % 10
		num /= 10
	}
	return res
}

