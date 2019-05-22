#
# @lc app=leetcode.cn id=102 lang=python3
#
# [102] 二叉树的层次遍历
#
# https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
#
# algorithms
# Medium (55.88%)
# Likes:    201
# Dislikes: 0
# Total Accepted:    25.9K
# Total Submissions: 46.3K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
#
# 例如:
# 给定二叉树: [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
#
# 返回其层次遍历结果：
#
# [
# ⁠ [3],
# ⁠ [9,20],
# ⁠ [15,7]
# ]
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
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []
        if root == None:
            return res
        queue = [root]
        while len(queue):
            temp = []
            arr = []
            for node in queue:
                if node.left != None:
                    temp.append(node.left)
                if node.right != None:
                    temp.append(node.right)
                arr.append(node.val)
            res.append(arr)
            queue = temp
        return res
