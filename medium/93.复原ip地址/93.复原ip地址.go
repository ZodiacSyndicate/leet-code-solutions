/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (43.20%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 18.3K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 * 
 * 示例:
 * 
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 * 
 */

 import (
	"strconv"
	"strings"
 )
 func restoreIpAddresses(s string) []string {
	var res []string
	backtrack(s, []string{}, &res)
	return res
}

func backtrack(str string, temp []string, res *[]string) {
	if len(temp) == 3 {
		if (len(str) > 1 && str[0] == '0') || len(str) == 0 {
			return
		}
		num, _ := strconv.Atoi(str)
		if len(str) > 0 && num <= 255 {
			temp = append(temp, str)
			*res = append(*res, strings.Join(temp, "."))
		}
		return
	}
	for i := 1; i < 4; i++ {
		if len(str) < i {
			return
		}
		s := str[0:i]
		num, _ := strconv.Atoi(s)
		if len(s) > 1 && (num > 255 || s[0] == '0') {
			return
		}
		arr := temp[:]
		arr = append(arr, s)
		backtrack(str[i:], arr, res)
	}
}

