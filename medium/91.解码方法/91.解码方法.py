#
# @lc app=leetcode.cn id=91 lang=python3
#
# [91] 解码方法
#
# https://leetcode-cn.com/problems/decode-ways/description/
#
# algorithms
# Medium (20.09%)
# Likes:    126
# Dislikes: 0
# Total Accepted:    7.6K
# Total Submissions: 37.4K
# Testcase Example:  '"12"'
#
# 一条包含字母 A-Z 的消息通过以下方式进行了编码：
#
# 'A' -> 1
# 'B' -> 2
# ...
# 'Z' -> 26
#
#
# 给定一个只包含数字的非空字符串，请计算解码方法的总数。
#
# 示例 1:
#
# 输入: "12"
# 输出: 2
# 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
#
#
# 示例 2:
#
# 输入: "226"
# 输出: 3
# 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
#
#
#


class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == '0':
            return 0
        dp = [1, 1]
        for i in range(len(s)):
            if s[i - 1] != '0':
                dp.append(0)
                num = int(s[i-1] + s[i]) if i > 0 else 0
                if num >= 1 and num <= 26:
                    dp[i + 1] = dp[i - 1] + dp[i] if s[i] != '0' else dp[i - 1]
                elif s[i] != '0':
                    dp[i + 1] = dp[i]
                else:
                    return 0
            elif s[i] != '0':
                dp.append(0)
                dp[i + 1] = dp[i]
            else:
                return 0
        return dp[len(s)]
