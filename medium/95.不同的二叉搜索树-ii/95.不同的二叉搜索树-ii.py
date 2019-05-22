#
# @lc app=leetcode.cn id=95 lang=python3
#
# [95] 不同的二叉搜索树 II
#
# https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
#
# algorithms
# Medium (54.81%)
# Likes:    132
# Dislikes: 0
# Total Accepted:    5.5K
# Total Submissions: 10.1K
# Testcase Example:  '3'
#
# 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
#
# 示例:
#
# 输入: 3
# 输出:
# [
# [1,null,3,2],
# [3,2,null,1],
# [3,1,null,null,2],
# [2,1,3],
# [1,null,2,null,3]
# ]
# 解释:
# 以上的输出对应以下 5 种不同结构的二叉搜索树：
#
# ⁠  1         3     3      2      1
# ⁠   \       /     /      / \      \
# ⁠    3     2     1      1   3      2
# ⁠   /     /       \                 \
# ⁠  2     1         2                 3
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
    def generateTrees(self, n):
        if n == 0:
            return []
        return self.buildTree(1, n)

    def buildTree(self, start, end):
        res = []
        if start > end:
            res.append(None)
            return res
        for i in range(start, end + 1):
            subLeft = self.buildTree(start, i - 1)
            subRight = self.buildTree(i + 1, end)
            for left in subLeft:
                for right in subRight:
                    node = TreeNode(i)
                    node.left = left
                    node.right = right
                    res.append(node)
        return res
