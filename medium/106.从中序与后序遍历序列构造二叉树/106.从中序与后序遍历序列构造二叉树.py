#
# @lc app=leetcode.cn id=106 lang=python3
#
# [106] 从中序与后序遍历序列构造二叉树
#
# https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
#
# algorithms
# Medium (60.75%)
# Likes:    79
# Dislikes: 0
# Total Accepted:    7.9K
# Total Submissions: 13K
# Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
#
# 根据一棵树的中序遍历与后序遍历构造二叉树。
#
# 注意:
# 你可以假设树中没有重复的元素。
#
# 例如，给出
#
# 中序遍历 inorder = [9,3,15,20,7]
# 后序遍历 postorder = [9,15,7,20,3]
#
# 返回如下的二叉树：
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
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
    def buildTree(self, inorder: List[int], postorder: List[int]) -> TreeNode:
        def build(il, ir, pl, pr):
            if il > ir or pl > pr:
                return None
            node = TreeNode(postorder[pr])
            index = 0
            for i in range(il, ir+1):
                if postorder[pr] == inorder[i]:
                    index = i
                    break
            rightCount = ir - index
            node.left = build(il, index - 1, pl, pr - rightCount - 1)
            node.right = build(index + 1, ir, pr - rightCount, pr - 1)
            return node
        return build(0, len(inorder) - 1, 0, len(postorder) - 1)
