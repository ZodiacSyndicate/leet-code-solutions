/*
 * @lc app=leetcode.cn id=21 lang=javascript
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.19%)
 * Total Accepted:    42.2K
 * Total Submissions: 80.8K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
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
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
  if (!l1 && !l2) return null
  let l3 = new ListNode(null)
  const head = l3

  while (l1 && l2) {
    if (l1.val <= l2.val) {
      l3.next = new ListNode(l1.val)
      l1 = l1.next
    } else {
      l3.next = new ListNode(l2.val)
      l2 = l2.next
    }
    l3 = l3.next
  }

  if (l1) l3.next = l1
  if (l2) l3.next = l2
  return head.next
}
