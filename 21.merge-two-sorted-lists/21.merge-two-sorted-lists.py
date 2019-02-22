#
# @lc app=leetcode.cn id=21 lang=python3
#
# [21] 合并两个有序链表
#
# https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
#
# algorithms
# Easy (52.19%)
# Total Accepted:    42.2K
# Total Submissions: 80.8K
# Testcase Example:  '[1,2,4]\n[1,3,4]'
#
# 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
#
# 示例：
#
# 输入：1->2->4, 1->3->4
# 输出：1->1->2->3->4->4
#
#
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: 'ListNode', l2: 'ListNode') -> 'ListNode':
        if not l1 and not l2:
            return None
        l3 = ListNode(0)
        head = l3
        while l1 and l2:
            newNode = ListNode(None)
            if l1.val <= l2.val:
                newNode.val = l1.val
                l1 = l1.next
            else:
                newNode.val = l2.val
                l2 = l2.next
            l3.next = newNode
            l3 = l3.next
        if l1:
            l3.next = l1
        if l2:
            l3.next = l2

        return head.next
