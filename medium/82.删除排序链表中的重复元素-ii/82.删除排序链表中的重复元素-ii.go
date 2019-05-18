/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (39.76%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 25K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -100000
	dummy.Next = head
	node, prev := dummy, dummy
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			for node != nil && node.Next != nil && node.Val == node.Next.Val {
				node = node.Next
			}
			prev.Next = node.Next
			node = node.Next
		} else {
			prev = node
			node = node.Next
		}
	}
	return dummy.Next
}

