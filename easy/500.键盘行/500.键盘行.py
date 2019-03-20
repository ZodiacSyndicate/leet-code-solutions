#
# @lc app=leetcode.cn id=500 lang=python3
#
# [500] 键盘行
#
# https://leetcode-cn.com/problems/keyboard-row/description/
#
# algorithms
# Easy (64.94%)
# Total Accepted:    5.7K
# Total Submissions: 8.8K
# Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
#
# 给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
#
#
#
#
#
#
#
# 示例：
#
# 输入: ["Hello", "Alaska", "Dad", "Peace"]
# 输出: ["Alaska", "Dad"]
#
#
#
#
# 注意：
#
#
# 你可以重复使用键盘上同一字符。
# 你可以假设输入的字符串将只包含字母。
#
#


class Solution:
    line1 = 'qwertyuiopQWERTYUIOP'
    line2 = 'asdfghjklASDFGHJKL'
    line3 = 'zxcvbnmZXCVBNM'

    def findWords(self, words: List[str]) -> List[str]:
        ans = []
        line = ''
        for word in words:
            if word[0] in self.line1:
                line = self.line1
            elif word[0] in self.line2:
                line = self.line2
            else:
                line = self.line3
            flag = True
            for s in word:
                if s not in line:
                    flag = False
                    break
            if flag:
                ans.append(word)
        return ans
