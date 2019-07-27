/*
 * @lc app=leetcode.cn id=7 lang=rust
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.69%)
 * Likes:    1179
 * Dislikes: 0
 * Total Accepted:    153.2K
 * Total Submissions: 468.6K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 * 
 * 示例 1:
 * 
 * 输入: 123
 * 输出: 321
 * 
 * 
 * 示例 2:
 * 
 * 输入: -123
 * 输出: -321
 * 
 * 
 * 示例 3:
 * 
 * 输入: 120
 * 输出: 21
 * 
 * 
 * 注意:
 * 
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 * 
 */
impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut res: i32 = 0;
        let mut x = x;
        while x != 0 {
            if let Some(n) = res.checked_mul(10) {
                res = n + x % 10;
                x /= 10;
            } else {
                return 0;
            }
        }
        res
    }
}

