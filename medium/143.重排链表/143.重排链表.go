/*
 * @lc app=leetcode.cn id=143 lang=golang
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
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	p := head
	l := []*ListNode{}
	for p != nil {
		l = append(l, p)
		p = p.Next
	}

	n := len(l)
	for i := 0; i < n/2; i++ {
		l[i].Next = l[n-i-1]
		l[n-i-1].Next = l[i+1]
	}
	if n > 0 {
		l[n/2].Next = nil
	}
}

