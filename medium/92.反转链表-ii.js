/*
 * @lc app=leetcode.cn id=92 lang=javascript
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (42.82%)
 * Likes:    136
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 23.5K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */
var reverseBetween = function(head, m, n) {
  const dummy = new ListNode(null)
  dummy.next = head
  let i = 1
  let node = dummy
  while (i < m) {
    node = node.next
    i++
  }
  let p = null
  let q
  while (i <= n) {
    const temp = node.next
    if (!q) {
      q = temp
    }
    node.next = node.next.next
    temp.next = p
    p = temp
    i++
  }
  q.next = node.next
  node.next = p
  return dummy.next
}
