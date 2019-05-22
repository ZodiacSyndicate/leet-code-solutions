#
# @lc app=leetcode.cn id=114 lang=python3
#
# [114] 二叉树展开为链表
#
# https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
#
# algorithms
# Medium (60.73%)
# Likes:    108
# Dislikes: 0
# Total Accepted:    6.8K
# Total Submissions: 11.2K
# Testcase Example:  '[1,2,5,3,4,null,6]'
#
# 给定一个二叉树，原地将它展开为链表。
#
# 例如，给定二叉树
#
# ⁠   1
# ⁠  / \
# ⁠ 2   5
# ⁠/ \   \
# 3   4   6
#
# 将其展开为：
#
# 1
# ⁠\
# ⁠ 2
# ⁠  \
# ⁠   3
# ⁠    \
# ⁠     4
# ⁠      \
# ⁠       5
# ⁠        \
# ⁠         6
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if not root:
            return
        if not root.left and not root.right:
            return
        if root.left:
            self.flatten(root.left)
        if root.right:
            self.flatten(root.right)
        if not root.left:
            return
        r = root.right
        l = root.left
        node = l
        while node and node.right:
            node = node.right
        node.right = r
        root.right = l
        root.left = None
