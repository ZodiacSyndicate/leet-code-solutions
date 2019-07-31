#
# @lc app=leetcode.cn id=146 lang=python3
#
# [146] LRU缓存机制
#
# https://leetcode-cn.com/problems/lru-cache/description/
#
# algorithms
# Medium (41.62%)
# Likes:    169
# Dislikes: 0
# Total Accepted:    10.3K
# Total Submissions: 24.6K
# Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
#
# 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
#
# 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
# 写入数据 put(key, value) -
# 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
#
# 进阶:
#
# 你是否可以在 O(1) 时间复杂度内完成这两种操作？
#
# 示例:
#
# LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
#
# cache.put(1, 1);
# cache.put(2, 2);
# cache.get(1);       // 返回  1
# cache.put(3, 3);    // 该操作会使得密钥 2 作废
# cache.get(2);       // 返回 -1 (未找到)
# cache.put(4, 4);    // 该操作会使得密钥 1 作废
# cache.get(1);       // 返回 -1 (未找到)
# cache.get(3);       // 返回  3
# cache.get(4);       // 返回  4
#
#
#


class ListNode:
    def __init__(self, val):
        self.val = val
        self.prev = self.next = self.key = None


class LRUCache:

    def __init__(self, capacity: int):
        self.map = {}
        self.head = None
        self.tail = None
        self.capacity = capacity
        self.length = 0

    def get(self, key: int) -> int:
        if self.map.get(key):
            node = self.map.get(key)
            if self.capacity == 1:
                return node.val
            if node == self.tail or self.head == self.tail:
                return node.val
            if node == self.head:
                self.head = self.head.next
                self.head.prev = None
                self.tail.next = node
                node.prev = self.tail
                node.next = None
                self.tail = node
                return node.val
            node.prev.next = node.next
            node.next.prev = node.prev
            self.tail.next = node
            node.prev = self.tail
            node.next = None
            self.tail = node
            return node.val
        return -1

    def put(self, key: int, value: int) -> None:
        if self.map.get(key):
            self.get(key)
            self.tail.val = value
            self.map[key] = self.tail
        else:
            node = ListNode(value)
            node.key = key
            if self.capacity == 1:
                self.map = {key: node}
                self.head = self.tail = node
                return
            if self.length == self.capacity:
                del self.map[self.head.key]
                self.map[key] = node
                self.head = self.head.next
                self.head.prev = None
                self.tail.next = node
                node.prev = self.tail
                self.tail = node
            else:
                self.length += 1
                self.map[key] = node
                if not self.head and not self.tail:
                    self.head = self.tail = node
                    return
                self.tail.next = node
                node.prev = self.tail
                self.tail = node
                # Your LRUCache object will be instantiated and called as such:
                # obj = LRUCache(capacity)
                # param_1 = obj.get(key)
                # obj.put(key,value)


cache = LRUCache(3)

cache.put(1, 1)
cache.put(2, 2)
cache.put(3, 3)
cache.put(4, 4)
cache.get(4)
cache.get(3)
cache.get(2)
cache.get(1)
cache.put(5, 5)
cache.get(1)
cache.get(2)
cache.get(3)
cache.get(4)
cache.get(5)
