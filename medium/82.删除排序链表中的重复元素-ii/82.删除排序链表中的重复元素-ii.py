#
# @lc app=leetcode.cn id=82 lang=python3
#
# [82] 删除排序链表中的重复元素 II
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        dummy = ListNode(-100000)
        dummy.next = head
        node, prev = dummy, dummy
        while node != None and node.next != None:
            if node.val == node.next.val:
                while node != None and node.next != None and node.val == node.next.val:
                    node = node.next
                prev.next = node.next
                node = node.next
            else:
                prev = node
                node = node.next

        return dummy.next

