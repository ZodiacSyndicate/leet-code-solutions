/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 *
 * https://leetcode-cn.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (20.09%)
 * Likes:    126
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 37.4K
 * Testcase Example:  '"12"'
 *
 * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
 * 
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 * 
 * 
 * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
 * 
 * 示例 1:
 * 
 * 输入: "12"
 * 输出: 2
 * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
 * 
 * 
 * 示例 2:
 * 
 * 输入: "226"
 * 输出: 3
 * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 * 
 * 
 */
import "strconv"
func numDecodings(s string) int {
  if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[0] = 1
	dp[1] = 1
	for i := range s {
		if i > 0 && s[i - 1] != '0' {
			num, _ := strconv.Atoi(string(s[i-1])+string(s[i]))
			if num >= 1 && num <= 26 {
				if s[i] != '0' {
					dp[i + 1] = dp[i - 1] + dp[i]
				} else {
					dp[i + 1] = dp[i - 1]
				} 
			} else if s[i] != '0' {
				dp[i + 1] = dp[i]
			}  else {
				return 0
			}
		} else if s[i] != '0' {
			dp[i + 1] = dp[i]
		} else {
			return 0
		}
	}
	return dp[len(s)]
}

