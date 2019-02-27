/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (37.90%)
 * Total Accepted:    22.4K
 * Total Submissions: 59.1K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 * 
 * 示例 1:
 * 
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "race a car"
 * 输出: false
 * 
 * 
 */
import "strings"
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	r := strings.Split(s, "")
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			continue
		}
		if s[i] < 97 || s[i] > 122 {
			r[i] = " "
		}
	}

	s = strings.Join(r, " ")
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	i := 0
	j := len(s)-1;
	for i <= j {
		if string(s[i]) != string(s[len(s)-i-1]) {
			return false
		}
		i++
		j--
	}
	return true
}

