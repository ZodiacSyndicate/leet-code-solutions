#
# @lc app=leetcode.cn id=105 lang=python3
#
# [105] 从前序与中序遍历序列构造二叉树
#
# https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
#
# algorithms
# Medium (57.84%)
# Likes:    158
# Dislikes: 0
# Total Accepted:    12.8K
# Total Submissions: 22.1K
# Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
#
# 根据一棵树的前序遍历与中序遍历构造二叉树。
#
# 注意:
# 你可以假设树中没有重复的元素。
#
# 例如，给出
#
# 前序遍历 preorder = [3,9,20,15,7]
# 中序遍历 inorder = [9,3,15,20,7]
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
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        def build(il, ir, pl, pr):
            if pl > pr or il > ir:
                return None
            node = TreeNode(preorder[pl])
            index = 0
            for i in range(il, ir + 1):
                if inorder[i] == preorder[pl]:
                    index = i
                    break
            leftCount = index - il
            node.left = build(il, index - 1, pl + 1, pl + leftCount)
            node.right = build(index + 1, ir, pl + leftCount + 1, pr)
            return node
        return build(0, len(inorder) - 1, 0, len(preorder) - 1)
