/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (46.14%)
 * Total Accepted:    15.9K
 * Total Submissions: 34.3K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
import "strconv"

func addBinary(a string, b string) string {
	res := ""
	flag := false
	for len(a) != 0 || len(b) != 0 {
		if len(a) == 0 {
			a = "0"
		}
		if len(b) == 0 {
			b = "0"
		}
		a1, err1 := strconv.Atoi(string(a[len(a)-1]))
		b1, err2 := strconv.Atoi(string(b[len(b)-1]))
		if err1 != nil {
			a1 = 0
		}
		if err2 != nil {
			b1 = 0
		}
		if flag {
			if a1+b1+1 < 2 {
				flag = false
			}
			res = strconv.Itoa((a1+b1+1)%2) + res
		} else {
			if a1+b1 == 2 {
				flag = true
			}
			res = strconv.Itoa((a1+b1)%2) + res
		}
		a, b = a[:len(a)-1], b[:len(b)-1]
	}
	if flag {
		res = "1" + res
	}
	return res
}
