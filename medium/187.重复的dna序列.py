#
# @lc app=leetcode.cn id=187 lang=python3
#
# [187] 重复的DNA序列
#
# https://leetcode-cn.com/problems/repeated-dna-sequences/description/
#
# algorithms
# Medium (42.75%)
# Likes:    30
# Dislikes: 0
# Total Accepted:    3.7K
# Total Submissions: 8.7K
# Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
#
# 所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA
# 中的重复序列有时会对研究非常有帮助。
#
# 编写一个函数来查找 DNA 分子中所有出现超多一次的10个字母长的序列（子串）。
#
# 示例:
#
# 输入: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
#
# 输出: ["AAAAACCCCC", "CCCCCAAAAA"]
#
#


class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        m = {}
        res = set()
        for i in range(len(s) - 9):
            str = s[i:i+10]
            if str in m:
                res.add(str)
            else:
                m[str] = 1
        return list(res)
