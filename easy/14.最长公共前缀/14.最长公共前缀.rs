/*
 * @lc app=leetcode.cn id=14 lang=rust
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.16%)
 * Likes:    628
 * Dislikes: 0
 * Total Accepted:    100.1K
 * Total Submissions: 293.1K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 示例 1:
 * 
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 * 
 * 
 * 说明:
 * 
 * 所有输入只包含小写字母 a-z 。
 * 
 */
use std::cmp;
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        if strs.is_empty() || strs[0] == "" {
            return "".to_string();
        }
        if strs.len() == 1 {
            return strs[0].clone();
        }
        let mut res = String::from("");
        let mut temp :u8 = 0;
        let mut ptr = 0;

        loop {
            for i in 0..strs.len() {
                if ptr >= strs[i].len() {
                    return res.to_string();
                } else if i == 0 {
                    temp = strs[i].as_bytes()[ptr];
                } else if strs[i].as_bytes()[ptr] != temp {
                    return res.to_string();
                }
            }
            let str = strs[0].as_bytes()[ptr];
            res.push(str as char);
            ptr += 1
        }
    }
}

