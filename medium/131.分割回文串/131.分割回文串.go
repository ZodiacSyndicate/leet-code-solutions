/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (62.32%)
 * Likes:    115
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 12.3K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */
func partition(s string) [][]string {
	rst := [][]string{}
	curr := []string{}
	rst = bt(rst, curr, s, 0)
	return rst
}

func bt(rst [][]string, curr []string, s string, l int) [][]string {
	if len(curr) > 0 && l >= len(s) {
		tmp := []string{}
		for _, v := range curr {
			tmp = append(tmp, v)
		}
		rst = append(rst, tmp)
	}
	for i := l; i < len(s); i++ {
		if palindrome(s[l : i+1]) {
			if l == i {
				curr = append(curr, string(s[i]))
			} else {
				curr = append(curr, s[l:i+1])
			}
			rst = bt(rst, curr, s, i+1)
			curr = curr[:len(curr)-1]
		}
	}
	return rst
}

func palindrome(s string) bool {
	begin := 0
	end := len(s) - 1
	for begin < end {
		if s[begin] == s[end] {
			begin++
			end--
		} else {
			return false
		}
	}
	return true
}

