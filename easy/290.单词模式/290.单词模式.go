/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词模式
 *
 * https://leetcode-cn.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (36.88%)
 * Total Accepted:    4.8K
 * Total Submissions: 12.9K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * 给定一种 pattern(模式) 和一个字符串 str ，判断 str 是否遵循相同的模式。
 *
 * 这里的遵循指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应模式。
 *
 * 示例1:
 *
 * 输入: pattern = "abba", str = "dog cat cat dog"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:pattern = "abba", str = "dog cat cat fish"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: pattern = "aaaa", str = "dog cat cat dog"
 * 输出: false
 *
 * 示例 4:
 *
 * 输入: pattern = "abba", str = "dog dog dog dog"
 * 输出: false
 *
 * 说明:
 * 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 *
 */
import "strings"

func wordPattern(pattern string, str string) bool {
	m1, m2 := map[string]string{}, map[string]string{}
	arr := strings.Split(str, " ")
	if len(pattern) != len(arr) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		p := string(pattern[i])
		s := arr[i]
		if _, ok := m1[p]; ok {
			if m1[p] != s {
				return false
			}
		} else {
			m1[p] = s
		}
		if _, ok := m2[s]; ok {
			if m2[s] != p {
				return false
			}
		} else {
			m2[s] = p
		}
	}
	return true
}

