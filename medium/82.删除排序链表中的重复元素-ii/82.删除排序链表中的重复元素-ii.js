/*
 * @lc app=leetcode.cn id=82 lang=javascript
 *
 * [82] 删除排序链表中的重复元素 II
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
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
  const dummyHead = new ListNode(-(1 << 20))
  dummyHead.next = head
  let node = dummyHead
  let prev = dummyHead
  while (node && node.next) {
    if (node.val === node.next.val) {
      while (node && node.next && node.val === node.next.val) {
        node = node.next
      }
      prev.next = node.next
      node = node.next
    } else {
      prev = node
      node = node.next
    }
  }
  return dummyHead.next
}
