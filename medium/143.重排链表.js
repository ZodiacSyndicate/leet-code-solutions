/*
 * @lc app=leetcode.cn id=143 lang=javascript
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (49.55%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    5.7K
 * Total Submissions: 11.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 *
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 *
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
 * @return {void} Do not return anything, modify head in-place instead.
 */
var reorderList = function(head) {
  if (!head || !head.next) return

  let p1 = head
  let p2 = head

  while (p2.next && p2.next.next) {
    p1 = p1.next
    p2 = p2.next.next
  }

  p2 = p1.next
  p1.next = null
  p1 = head

  let head2 = p2
  let next2
  while (p2.next) {
    next2 = p2.next
    p2.next = next2.next
    next2.next = head2
    head2 = next2
  }
  p2 = head2

  let next1

  while (p2) {
    next1 = p1.next
    next2 = p2.next

    p1.next = p2
    p2.next = next1

    p1 = next1
    p2 = next2
  }
}
