#
# @lc app=leetcode.cn id=133 lang=python3
#
# [133] 克隆图
#
# https://leetcode-cn.com/problems/clone-graph/description/
#
# algorithms
# Medium (30.03%)
# Likes:    43
# Dislikes: 0
# Total Accepted:    3.5K
# Total Submissions: 11.4K
# Testcase Example:  '{"$id":"1","neighbors":[{"$id":"2","neighbors":[{"$ref":"1"},{"$id":"3","neighbors":[{"$ref":"2"},{"$id":"4","neighbors":[{"$ref":"3"},{"$ref":"1"}],"val":4}],"val":3}],"val":2},{"$ref":"4"}],"val":1}'
#
# 给定无向连通图中一个节点的引用，返回该图的深拷贝（克隆）。图中的每个节点都包含它的值 val（Int） 和其邻居的列表（list[Node]）。
#
# 示例：
#
#
#
# 输入：
#
# {"$id":"1","neighbors":[{"$id":"2","neighbors":[{"$ref":"1"},{"$id":"3","neighbors":[{"$ref":"2"},{"$id":"4","neighbors":[{"$ref":"3"},{"$ref":"1"}],"val":4}],"val":3}],"val":2},{"$ref":"4"}],"val":1}
#
# 解释：
# 节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
# 节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
# 节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
# 节点 4 的值是 4，它有两个邻居：节点 1 和 3 。
#
#
#
#
# 提示：
#
#
# 节点数介于 1 到 100 之间。
# 无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
# 由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p 的邻居。
# 必须将给定节点的拷贝作为对克隆图的引用返回。
#
#
#
"""
# Definition for a Node.
class Node:
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors
"""


class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        m = {}
        return self.dfs(node, m)

    def dfs(self, node, m):
        if not node:
            return node
        v = Node(node.val, [])
        m[v.val] = v

        for w in node.neighbors:
            if w.val in m:
                v.neighbors.append(m[w.val])
            else:
                v.neighbors.append(self.dfs(w, m))

        return v
