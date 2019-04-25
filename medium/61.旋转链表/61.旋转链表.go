/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	cursor, length := head, 1
	var tail *ListNode
	for cursor.Next != nil {
		cursor = cursor.Next
		length++
	}
	steps := length - (k % length)
	tail = cursor
	cursor.Next = head
	cursor = head
	for i := 0; i < steps; i++ {
		cursor = cursor.Next
		tail = tail.Next
	}
	tail.Next = nil
	return cursor
}

