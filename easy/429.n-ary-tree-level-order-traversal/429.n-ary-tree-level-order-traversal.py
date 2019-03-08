#
# @lc app=leetcode.cn id=429 lang=python3
#
# [429] N-ary Tree Level Order Traversal
#
# https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/description/
#
# algorithms
# Easy (59.24%)
# Total Accepted:    3.5K
# Total Submissions: 6K
# Testcase Example:  '{"$id":"1","children":[{"$id":"2","children":[{"$id":"5","children":[],"val":5},{"$id":"6","children":[],"val":6}],"val":3},{"$id":"3","children":[],"val":2},{"$id":"4","children":[],"val":4}],"val":1}'
#
# 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
#
# 例如，给定一个 3叉树 :
#
#
#
#
#
#
#
# 返回其层序遍历:
#
# [
# ⁠    [1],
# ⁠    [3,2,4],
# ⁠    [5,6]
# ]
#
#
#
#
# 说明:
#
#
# 树的深度不会超过 1000。
# 树的节点总数不会超过 5000。
#
#
"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""


class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        if root is None:
            return []
        ans = []
        cur = [root]
        next = []
        while len(cur) != 0:
            arr = []
            for node in cur:
                if node:
                    arr.append(node.val)
                    for child in node.children:
                        next.append(child)
            ans.append(arr)
            cur = next
            next = []
        return ans
