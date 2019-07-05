/*
 * @lc app=leetcode.cn id=146 lang=javascript
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (41.62%)
 * Likes:    169
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 24.6K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 * 示例:
 *
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 )
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *
 */

class ListNode {
  constructor(val) {
    this.val = val
    this.prev = this.next = this.key = null
  }
}
/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
  this.map = new Map()
  this.head = null
  this.tail = null
  this.capacity = capacity
  this.length = 0
}

/**
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
  if (this.map.has(key)) {
    const node = this.map.get(key)
    if (this.capacity === 1) return node.val
    if (node === this.tail || this.head === this.tail) return node.val
    if (node === this.head) {
      this.head = this.head.next
      this.head.prev = null
      this.tail.next = node
      node.prev = this.tail
      node.next = null
      this.tail = node
      return node.val
    }
    const { prev, next, val } = node
    prev.next = next
    next.prev = prev
    this.tail.next = node
    node.prev = this.tail
    node.next = null
    this.tail = node
    return val
  }
  return -1
}

/**
 * @param {number} key
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
  if (this.map.has(key)) {
    this.get(key)
    this.tail.val = value
    this.map.set(key, this.tail)
  } else {
    const node = new ListNode(value)
    node.key = key
    if (this.capacity === 1) {
      this.map = new Map([[key, node]])
      this.head = this.tail = node
      return
    }
    if (this.length === this.capacity) {
      this.map.delete(this.head.key)
      this.map.set(key, node)
      this.head = this.head.next
      this.head.prev = null
      this.tail.next = node
      node.prev = this.tail
      this.tail = node
    } else {
      this.length++
      this.map.set(key, node)
      if (!this.head && !this.tail) {
        this.head = this.tail = node
        return
      }
      this.tail.next = node
      node.prev = this.tail
      this.tail = node
    }
  }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
