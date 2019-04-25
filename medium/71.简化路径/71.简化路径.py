#
# @lc app=leetcode.cn id=71 lang=python3
#
# [71] 简化路径
#


class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = ['']
        arr = path.split('/')
        for s in arr:
            if s == '' or s == '.':
                continue
            elif s == '..':
                stack.pop()
                if len(stack) == 0:
                    stack.append('')
            else:
                stack.append(s)
        return '/' if len(stack) == 1 else '/'.join(stack)
