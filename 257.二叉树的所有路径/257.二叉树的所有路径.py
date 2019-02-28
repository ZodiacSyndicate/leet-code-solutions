#
# @lc app=leetcode.cn id=257 lang=python3
#
# [257] 二叉树的所有路径
#
# https://leetcode-cn.com/problems/binary-tree-paths/description/
#
# algorithms
# Easy (56.91%)
# Total Accepted:    5.9K
# Total Submissions: 10.4K
# Testcase Example:  '[1,2,3,null,5]'
#
# 给定一个二叉树，返回所有从根节点到叶子节点的路径。
# 
# 说明: 叶子节点是指没有子节点的节点。
# 
# 示例:
# 
# 输入:
# 
# ⁠  1
# ⁠/   \
# 2     3
# ⁠\
# ⁠ 5
# 
# 输出: ["1->2->5", "1->3"]
# 
# 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
# 
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def binaryTreePaths(self, root: TreeNode) -> List[str]:
        if root is None:
            return []
        res = []
        stack = [(str(root.val), root)]
        while stack:
            cur, node = stack.pop()
            if node:
                if not node.left and not node.right:
                    res.append(cur)
                if node.left:
                    stack.append((cur + '->' + str(node.left.val), node.left))
                if node.right:
                    stack.append((cur + '->' + str(node.right.val), node.right))
        return res
