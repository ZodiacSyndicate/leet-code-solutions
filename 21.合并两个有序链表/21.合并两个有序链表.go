/*
 * @lc app=leetcode.cn id=21 lang=golang
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
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l3 := new(ListNode)
	l3.Val = 0
	head := l3

	for l1 != nil && l2 != nil {
		newNode := new(ListNode)
		if l1.Val <= l2.Val {
			newNode.Val = l1.Val
			l1 = l1.Next
		} else {
			newNode.Val = l2.Val
			l2 = l2.Next
		}
		l3.Next = newNode
		l3 = l3.Next
	}

	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return head.Next
}
