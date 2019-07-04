#
# @lc app=leetcode.cn id=139 lang=python3
#
# [139] 单词拆分
#
# https://leetcode-cn.com/problems/word-break/description/
#
# algorithms
# Medium (41.03%)
# Likes:    119
# Dislikes: 0
# Total Accepted:    10.3K
# Total Submissions: 24.9K
# Testcase Example:  '"leetcode"\n["leet","code"]'
#
# 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
#
# 说明：
#
#
# 拆分时可以重复使用字典中的单词。
# 你可以假设字典中没有重复的单词。
#
#
# 示例 1：
#
# 输入: s = "leetcode", wordDict = ["leet", "code"]
# 输出: true
# 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
#
#
# 示例 2：
#
# 输入: s = "applepenapple", wordDict = ["apple", "pen"]
# 输出: true
# 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
# 注意你可以重复使用字典中的单词。
#
#
# 示例 3：
#
# 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
# 输出: false
#
#
#


class Solution:
    def wordBreak(self, s: str, wordDict) -> bool:
        memo = [None] * len(s)
        words = set(wordDict)
        return self.backtrack(s, 0, memo, words)

    def backtrack(self, s, start, memo, words):
        if start == len(s):
            return True
        if memo[start] != None:
            return memo[start]
        for end in range(start + 1, len(s) + 1):
            if s[start:end] in words and self.backtrack(s, end, memo, words):
                memo[start] = True
                return True
        memo[start] = False
        return False
