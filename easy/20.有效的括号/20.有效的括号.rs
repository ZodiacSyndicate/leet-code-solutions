/*
 * @lc app=leetcode.cn id=20 lang=rust
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.94%)
 * Likes:    937
 * Dislikes: 0
 * Total Accepted:    98K
 * Total Submissions: 251.6K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 */
use std::collections::HashMap;
impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut map: HashMap<char, char> = HashMap::new();
        map.insert('[', ']');
        map.insert('(', ')');
        map.insert('{', '}');
        let mut stack: Vec<char> = vec![];
        for c in s.chars() {
            if let Some(s) = map.get(&c) {
                stack.push(c);
            } else {
                let last = stack.last();
                if last.is_some() && map.get(&last.unwrap()).is_some() && *map.get(&last.unwrap()).unwrap() == c {
                    stack.pop();
                } else {
                    stack.push(c);
                }
            }
        }
        stack.is_empty()
    }
}

