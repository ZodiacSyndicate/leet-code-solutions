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
	return backtrack(s, []string{}, &[][]string{})
}

func backtrack(str string, arr []string, res *[][]string) [][]string {
	if isPalidrome(str) {
		temp := arr[:]
		temp = append(temp, str)
		*res = append(*res, temp)
	}
	for i := 1; i < len(str); i++ {
		temp := arr[:]
		s := str[0:i]
		if !isPalidrome(s) {
			continue
		}
		temp = append(temp, s)
		backtrack(str[i:], temp, res)
	}
	return *res
}

func isPalidrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

