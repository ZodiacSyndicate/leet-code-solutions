#
# @lc app=leetcode.cn id=38 lang=python3
#
# [38] 报数
#
# https://leetcode-cn.com/problems/count-and-say/description/
#
# algorithms
# Easy (47.88%)
# Total Accepted:    20.2K
# Total Submissions: 42.1K
# Testcase Example:  '1'
#
# 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
#
# 1.     1
# 2.     11
# 3.     21
# 4.     1211
# 5.     111221
#
#
# 1 被读作  "one 1"  ("一个一") , 即 11。
# 11 被读作 "two 1s" ("两个一"）, 即 21。
# 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
#
# 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
#
# 注意：整数顺序将表示为一个字符串。
#
#
#
# 示例 1:
#
# 输入: 1
# 输出: "1"
#
#
# 示例 2:
#
# 输入: 4
# 输出: "1211"
#
#
#
class Solution:
    def countAndSay(self, n: 'int') -> 'str':
        res = '1'
        while n > 1:
            s = ''
            count = 1
            for i in range(len(res)):
                if i + 1 < len(res) and res[i] == res[i + 1]:
                    count += 1
                else:
                    s += str(count) + res[i]
                    count = 1
            res = s
            n -= 1
        return res
