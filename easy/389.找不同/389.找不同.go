/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 *
 * https://leetcode-cn.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (54.04%)
 * Total Accepted:    6.5K
 * Total Submissions: 12K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 * 给定两个字符串 s 和 t，它们只包含小写字母。
 *
 * 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *
 * 请找出在 t 中被添加的字母。
 *
 *
 *
 * 示例:
 *
 * 输入：
 * s = "abcd"
 * t = "abcde"
 *
 * 输出：
 * e
 *
 * 解释：
 * 'e' 是那个被添加的字母。
 *
 *
 */

func findTheDifference(s string, t string) byte {
	m := map[byte]int{}
	var res byte
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; !ok {
			res = t[i]
		} else {
			m[t[i]]--
			if m[t[i]] < 0 {
				res = t[i]
			}
		}
	}
	return res
}

