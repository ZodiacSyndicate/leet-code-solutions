#
# @lc app=leetcode.cn id=43 lang=python3
#
# [43] 字符串相乘
#
# https://leetcode-cn.com/problems/multiply-strings/description/
#
# algorithms
# Medium (37.97%)
# Total Accepted:    12.2K
# Total Submissions: 32.3K
# Testcase Example:  '"2"\n"3"'
#
# 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
# 
# 示例 1:
# 
# 输入: num1 = "2", num2 = "3"
# 输出: "6"
# 
# 示例 2:
# 
# 输入: num1 = "123", num2 = "456"
# 输出: "56088"
# 
# 说明：
# 
# 
# num1 和 num2 的长度小于110。
# num1 和 num2 只包含数字 0-9。
# num1 和 num2 均不以零开头，除非是数字 0 本身。
# 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
# 
# 
#
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        n1, n2 = len(num1) - 1, len(num2) - 1
        if n1 < 0 or n2 < 0:
            return ''
        mul = [0] * (n1 + n2 + 2)
        for i in range(n1, -1, -1):
            for j in range(n2, -1, -1):
                bitMul = (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))
                bitMul += mul[i + j + 1]
                mul[i + j] += bitMul // 10
                mul[i + j + 1] = bitMul % 10
        
        i = 0
        result = ''
        while i < len(mul) - 1 and mul[i] == 0:
            i += 1
        for n in mul[i:]:
            result += str(n)
        return result

