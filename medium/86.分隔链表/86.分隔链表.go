/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (46.68%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 16.9K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 * 
 * 你应当保留两个分区中每个节点的初始相对位置。
 * 
 * 示例:
 * 
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
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

func partition(head *ListNode, x int) *ListNode {
	smallHead := new(ListNode)
	bigHead := new(ListNode)
	res, big := smallHead, bigHead
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			bigHead.Next = head
			bigHead = bigHead.Next
		}
		head = head.Next
	}
	smallHead.Next = big.Next
	bigHead.Next = nil
	return res.Next
}

