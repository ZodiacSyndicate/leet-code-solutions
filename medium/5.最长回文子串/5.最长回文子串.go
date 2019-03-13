/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (24.64%)
 * Total Accepted:    40.1K
 * Total Submissions: 162.6K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
import "math"

func longestPalindrome(s string) string {
	arr := []rune{'$', '#'}
	for _, n := range s {
		arr = append(arr, n)
		arr = append(arr, '#')
	}
	arr = append(arr, '%')

	ans := ""
	maxPoint := 0
	rightMax := 0
	middle := 0
	lens := make([]int, len(arr))
	for i, _ := range arr {
		if len(arr)-1-i <= lens[maxPoint] {
			break
		}

		lens[i] = 1

		if rightMax > i {
			lens[i] = int(math.Min(float64(rightMax-1), float64(lens[2*middle-i])))
		}

		for i-lens[i] >= 0 && arr[i-lens[i]] == arr[i+lens[i]] {
			lens[i]++
		}

		if lens[i]+1 > lens[maxPoint] {
			maxPoint = i
		}
	}

	for i := maxPoint - (lens[maxPoint] - 2); i < maxPoint+lens[maxPoint]-1; i += 2 {
		ans += string(arr[i])
	}

	return ans
}

