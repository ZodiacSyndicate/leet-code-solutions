/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (44.79%)
 * Total Accepted:    2.5K
 * Total Submissions: 5.4K
 * Testcase Example:  '0'
 *
 * 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
 *
 * 每个 LED 代表一个 0 或 1，最低位在右侧。
 *
 *
 *
 * 例如，上面的二进制手表读取 “3:25”。
 *
 * 给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
 *
 * 案例:
 *
 *
 * 输入: n = 1
 * 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16",
 * "0:32"]
 *
 *
 *
 * 注意事项:
 *
 *
 * 输出的顺序没有要求。
 * 小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
 * 分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
 *
 *
 */
import "strconv"

func readBinaryWatch(num int) []string {
	ans := []string{}
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			count := 0
			time := (h << 6) | m
			for time != 0 {
				time &= time - 1
				count++
			}
			if count == num {
				hStr := strconv.Itoa(h)
				var mStr string
				if m > 9 {
					mStr = strconv.Itoa(m)
				} else {
					mStr = "0" + strconv.Itoa(m)
				}
				ans = append(ans, hStr+":"+mStr)
			}
		}
	}
	return ans
}

