/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (32.31%)
 * Total Accepted:    87.9K
 * Total Submissions: 268.4K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	head.Val = 0
	head.Next = nil
	node := head
	carry := false
	for l1 != nil && l2 != nil {
		var val int
		if carry {
			val = l1.Val + l2.Val + 1
			if val < 10 {
				carry = false
			}
		} else {
			val = l1.Val + l2.Val
			if val > 9 {
				carry = true
			}
		}
		val %= 10
		node.Next = new(ListNode)
		node.Next.Val = val
		node.Next.Next = nil
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		if carry {
			if l1.Val+1 < 10 {
				carry = false
			}
			node.Next = new(ListNode)
			node.Next.Val = (l1.Val + 1) % 10
			node.Next.Next = nil
			node = node.Next
			l1 = l1.Next
		} else {
			node.Next = l1
			break
		}
	}
	for l2 != nil {
		if carry {
			if l2.Val+1 < 10 {
				carry = false
			}
			node.Next = new(ListNode)
			node.Next.Val = (l2.Val + 1) % 10
			node.Next.Next = nil
			node = node.Next
			l2 = l2.Next
		} else {
			node.Next = l2
			break
		}
	}
	if carry && l1 == nil && l2 == nil {
		node.Next = new(ListNode)
		node.Next.Val = 1
		node.Next.Next = nil
	}
	return head.Next
}

