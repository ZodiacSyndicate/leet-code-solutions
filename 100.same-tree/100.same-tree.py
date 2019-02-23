#
# @lc app=leetcode.cn id=100 lang=python3
#
# [100] 相同的树
#
# https://leetcode-cn.com/problems/same-tree/description/
#
# algorithms
# Easy (50.91%)
# Total Accepted:    14.1K
# Total Submissions: 27.7K
# Testcase Example:  '[1,2,3]\n[1,2,3]'
#
# 给定两个二叉树，编写一个函数来检验它们是否相同。
#
# 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
#
# 示例 1:
#
# 输入:       1         1
# ⁠         / \       / \
# ⁠        2   3     2   3
#
# ⁠       [1,2,3],   [1,2,3]
#
# 输出: true
#
# 示例 2:
#
# 输入:      1          1
# ⁠         /           \
# ⁠        2             2
#
# ⁠       [1,2],     [1,null,2]
#
# 输出: false
#
#
# 示例 3:
#
# 输入:       1         1
# ⁠         / \       / \
# ⁠        2   1     1   2
#
# ⁠       [1,2,1],   [1,1,2]
#
# 输出: false
#
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isSameTree(self, p: 'TreeNode', q: 'TreeNode') -> 'bool':
        if p and q:
            return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        if not p and not q:
            return True
        return False
