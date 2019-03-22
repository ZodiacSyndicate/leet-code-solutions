#
# @lc app=leetcode.cn id=17 lang=python3
#
# [17] 电话号码的字母组合
#
# https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
#
# algorithms
# Medium (47.98%)
# Total Accepted:    18.1K
# Total Submissions: 37.6K
# Testcase Example:  '"23"'
#
# 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
#
# 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
#
#
#
# 示例:
#
# 输入："23"
# 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
#
#
# 说明:
# 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
#
#


class Solution:
    arr = {
        '2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl', '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'
    }

    def letterCombinations(self, digits: str) -> List[str]:
        ans = []
        if digits == '':
            return ans
        self.generate(digits, '', ans, 0)
        return ans

    def generate(self, digits, s, ans, k):
        if len(s) == len(digits):
            ans.append(s)
            return
        temp = self.arr[digits[k]]
        for i in range(len(temp)):
            s += temp[i]
            self.generate(digits, s, ans, k + 1)
            s = s[0:len(s) - 1]
        return
