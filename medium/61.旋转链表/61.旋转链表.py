#
# @lc app=leetcode.cn id=61 lang=python3
#
# [61] 旋转链表
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if head is None or k == 0:
            return head
        cursor, tail, length = head, None, 1
        while cursor.next:
            cursor = cursor.next
            length += 1
        steps = length - (k % length)
        tail = cursor
        cursor.next = head
        cursor = head
        for i in range(steps):
            cursor = cursor.next
            tail = tail.next
        tail.next = None
        return cursor
