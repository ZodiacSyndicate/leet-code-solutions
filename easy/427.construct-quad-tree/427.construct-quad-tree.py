#
# @lc app=leetcode.cn id=427 lang=python3
#
# [427] Construct Quad Tree
#
# https://leetcode-cn.com/problems/construct-quad-tree/description/
#
# algorithms
# Easy (53.25%)
# Total Accepted:    852
# Total Submissions: 1.6K
# Testcase Example:  '[[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]'
#
# 我们想要使用一棵四叉树来储存一个 N x N 的布尔值网络。网络中每一格的值只会是真或假。树的根结点代表整个网络。对于每个结点,
# 它将被分等成四个孩子结点直到这个区域内的值都是相同的.
#
# 每个结点还有另外两个布尔变量: isLeaf 和 val。isLeaf 当这个节点是一个叶子结点时为真。val 变量储存叶子结点所代表的区域的值。
#
# 你的任务是使用一个四叉树表示给定的网络。下面的例子将有助于你理解这个问题：
#
# 给定下面这个8 x 8 网络，我们将这样建立一个对应的四叉树：
#
#
#
# 由上文的定义，它能被这样分割：
#
#
#
#
#
# 对应的四叉树应该像下面这样，每个结点由一对 (isLeaf, val) 所代表.
#
# 对于非叶子结点，val 可以是任意的，所以使用 * 代替。
#
#
#
# 提示：
#
#
# N 将小于 1000 且确保是 2 的整次幂。
# 如果你想了解更多关于四叉树的知识，你可以参考这个 wiki 页面。
#
#
#
"""
# Definition for a QuadTree node.
class Node:
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight
"""


class Solution:
    def construct(self, grid: List[List[int]]) -> 'Node':
        return self.dfs(grid, 0, len(grid), 0, len(grid))

    def dfs(self, grid, left, right, top, bottom):
        root = None
        key = grid[top][left]
        for i in range(top, bottom):
            for j in range(left, right):
                if grid[i][j] != key:
                    topLeft = self.dfs(
                        grid, left, (left + right) // 2, top, (top + bottom) // 2)
                    topRight = self.dfs(
                        grid, (left + right) // 2, right, top, (top + bottom) // 2)
                    bottomLeft = self.dfs(
                        grid, left, (left + right) // 2, (top + bottom) // 2, bottom)
                    bottomRight = self.dfs(
                        grid, (left + right) // 2, right, (top + bottom) // 2, bottom)
                    root = Node(False, False, topLeft, topRight,
                                bottomLeft, bottomRight)
                    return root
        root = Node(True if key == 1 else False, True, None, None, None, None)
        return root
