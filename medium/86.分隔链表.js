/*
 * @lc app=leetcode.cn id=86 lang=javascript
 *
 * [86] 分隔链表
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
 * @param {number} x
 * @return {ListNode}
 */
var partition = function(head, x) {
  let smallHead = new ListNode(null)
  let bigHead = new ListNode(null)
  const res = smallHead
  const big = bigHead
  while (head) {
    if (head.val < x) {
      smallHead.next = head
      smallHead = smallHead.next
    } else {
      bigHead.next = head
      bigHead = bigHead.next
    }
    head = head.next
  }
  smallHead.next = big.next
  bigHead.next = null
  return res.next
}
