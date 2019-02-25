#
# @lc app=leetcode.cn id=111 lang=python3
#
# [111] 二叉树的最小深度
#
# https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
#
# algorithms
# Easy (36.80%)
# Total Accepted:    10.7K
# Total Submissions: 29.1K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，找出其最小深度。
#
# 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
#
# 说明: 叶子节点是指没有子节点的节点。
#
# 示例:
#
# 给定二叉树 [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
# 返回它的最小深度  2.
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

import sys


class Solution:
    def minDepth(self, root: 'TreeNode') -> 'int':
        if not root:
            return 0
        stack = [(1, root)]
        res = sys.maxsize
        while len(stack):
            cur, node = stack.pop()
            if node:
                stack.append((cur + 1, node.left))
                stack.append((cur + 1, node.right))
                if not node.left and not node.right:
                    res = min(res, cur)
        return res
