/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode-cn.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (51.51%)
 * Total Accepted:    11.4K
 * Total Submissions: 22K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数是不是“快乐数”。
 *
 * 一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到
 * 1。如果可以变为 1，那么这个数就是快乐数。
 *
 * 示例:
 *
 * 输入: 19
 * 输出: true
 * 解释:
 * 12 + 92 = 82
 * 82 + 22 = 68
 * 62 + 82 = 100
 * 12 + 02 + 02 = 1
 *
 *
 */
func isHappy(n int) bool {
	m := map[int]int{}
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = 1
		cur := 0
		for n != 0 {
			cur += (n % 10) * (n % 10)
			n /= 10
		}
		n = cur
	}
	return true
}

