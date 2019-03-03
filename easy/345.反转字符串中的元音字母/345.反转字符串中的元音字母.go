/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (45.92%)
 * Total Accepted:    6.6K
 * Total Submissions: 14.4K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 * 示例 1:
 *
 * 输入: "hello"
 * 输出: "holle"
 *
 *
 * 示例 2:
 *
 * 输入: "leetcode"
 * 输出: "leotcede"
 *
 * 说明:
 * 元音字母不包含字母"y"。
 *
 */
import "strings"

func reverseVowels(s string) string {
	vowels := "AEIOUaeiou"
	arr := []string{}
	for i := 0; i < len(s); i++ {
		arr = append(arr, string(s[i]))
	}
	l := 0
	r := len(arr) - 1
	for l < r {
		for strings.Contains(vowels, arr[l]) == false && l < r {
			l++
		}
		for strings.Contains(vowels, arr[r]) == false && l < r {
			r--
		}
		if l < r && arr[l] != arr[r] {
			arr[l], arr[r] = arr[r], arr[l]
		}
		l++
		r--
	}
	return strings.Join(arr, "")
}

