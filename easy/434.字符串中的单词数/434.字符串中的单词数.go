/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (28.91%)
 * Total Accepted:    3.9K
 * Total Submissions: 13.4K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 * 
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 * 
 * 示例:
 * 
 * 输入: "Hello, my name is John"
 * 输出: 5
 * 
 * 
 */
func countSegments(s string) int {
	flag := true
	count := 0
	for _, n := range s {
		if flag {
			if n != ' ' {
				flag = false
			}
		} else {
			if n == ' ' {
				count ++
				flag = true
			}
		}
	}
	if flag {
		return count
	}
	return count + 1
}

