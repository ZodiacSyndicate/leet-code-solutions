/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (57.80%)
 * Total Accepted:    10.2K
 * Total Submissions: 17.5K
 * Testcase Example:  '1'
 *
 * 写一个程序，输出从 1 到 n 数字的字符串表示。
 *
 * 1. 如果 n 是3的倍数，输出“Fizz”；
 *
 * 2. 如果 n 是5的倍数，输出“Buzz”；
 *
 * 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 *
 * 示例：
 *
 * n = 15,
 *
 * 返回:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 *
 *
 */
import "strconv"

func fizzBuzz(n int) []string {
	ans := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

