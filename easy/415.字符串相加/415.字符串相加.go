import "strconv"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (42.83%)
 * Total Accepted:    4.9K
 * Total Submissions: 11.3K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1)-1, len(num2)-1
	carry := 0
	ans := ""
	for l1 >= 0 || l2 >= 0 {
		n1, n2 := 0, 0
		if l1 >= 0 {
			n1, _ = strconv.Atoi(string(num1[l1]))
		}
		if l2 >= 0 {
			n2, _ = strconv.Atoi(string(num2[l2]))
		}
		val := carry + n1 + n2
		carry = val / 10
		val %= 10
		ans = strconv.Itoa(val) + ans
		l1--
		l2--
	}
	if carry == 0 {
		return ans
	}
	return "1" + ans
}

