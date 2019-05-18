#
# @lc app=leetcode.cn id=93 lang=python3
#
# [93] 复原IP地址
#
# https://leetcode-cn.com/problems/restore-ip-addresses/description/
#
# algorithms
# Medium (43.20%)
# Likes:    94
# Dislikes: 0
# Total Accepted:    8.1K
# Total Submissions: 18.3K
# Testcase Example:  '"25525511135"'
#
# 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
#
# 示例:
#
# 输入: "25525511135"
# 输出: ["255.255.11.135", "255.255.111.35"]
#
#


class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []

        def backtrack(string, temp, k):
            if k == 3:
                if len(string) > 1 and string[0] == '0':
                    return
                if len(string) <= 3 and len(string) > 0 and int(string) <= 255:
                    temp.append(string)
                    res.append('.'.join(temp))
                return
            for i in range(1, 4):
                a = string[:i]
                if len(a) > 1 and int(a) > 255:
                    return
                if len(a) > 1 and a[0] == '0':
                    return
                arr = temp[:]
                arr.append(a)
                backtrack(string[i:], arr, k + 1)
        backtrack(s, [], 0)
        return res
