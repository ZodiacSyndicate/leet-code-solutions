#
# @lc app=leetcode.cn id=103 lang=python3
#
# [103] 二叉树的锯齿形层次遍历
#
# https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
#
# algorithms
# Medium (49.29%)
# Likes:    61
# Dislikes: 0
# Total Accepted:    10.7K
# Total Submissions: 21.7K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
#
# 例如：
# 给定二叉树 [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
#
# 返回锯齿形层次遍历如下：
#
# [
# ⁠ [3],
# ⁠ [20,9],
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
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []
        if root is None:
            return res
        queue = [root]
        flag = True
        while len(queue):
            temp = []
            arr = []
            for node in queue:
                if node.left != None:
                    temp.append(node.left)
                if node.right != None:
                    temp.append(node.right)
            if not flag:
                queue.reverse()
            for node in queue:
                arr.append(node.val)
            res.append(arr)
            queue = temp
            flag = not flag
        return res
