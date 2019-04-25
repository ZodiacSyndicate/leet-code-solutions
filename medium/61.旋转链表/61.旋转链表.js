/*
 * @lc app=leetcode.cn id=61 lang=javascript
 *
 * [61] 旋转链表
 */
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function(head, k) {
  if (!head || k === 0) return head
  let cursor = head
  let tail = null
  let len = 1
  while (cursor.next) {
    cursor = cursor.next
    len++
  }
  let steps = len - (k % len)
  tail = cursor
  cursor.next = head
  cursor = head
  for (let i = 0; i < steps; i++) {
    cursor = cursor.next
    tail = tail.next
  }
  tail.next = null
  return cursor
}
