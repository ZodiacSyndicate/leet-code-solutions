#
# @lc app=leetcode.cn id=138 lang=python3
#
# [138] 复制带随机指针的链表
#
# https://leetcode-cn.com/problems/copy-list-with-random-pointer/description/
#
# algorithms
# Medium (30.36%)
# Likes:    100
# Dislikes: 0
# Total Accepted:    7K
# Total Submissions: 22.6K
# Testcase Example:  '{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}'
#
# 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
#
# 要求返回这个链表的深拷贝。 
#
#
#
# 示例：
#
#
#
# 输入：
#
# {"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
#
# 解释：
# 节点 1 的值是 1，它的下一个指针和随机指针都指向节点 2 。
# 节点 2 的值是 2，它的下一个指针指向 null，随机指针指向它自己。
#
#
#
#
# 提示：
#
#
# 你必须返回给定头的拷贝作为对克隆列表的引用。
#
#
#
"""
# Definition for a Node.
class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random
"""


class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        return self.copy(head, {})

    def copy(self, head, m):
        if not head:
            return head
        node = Node(head.val, None, None)
        m[node.val] = node
        if head.next and head.next.val in m:
            node.next = m[head.next.val]
        else:
            node.next = self.copy(head.next, m)
        if head.random and head.random.val in m:
            node.random = m[head.random.val]
        else:
            node.random = self.copy(head.random, m)
        return node
