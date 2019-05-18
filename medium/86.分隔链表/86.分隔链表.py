#
# @lc app=leetcode.cn id=86 lang=python3
#
# [86] 分隔链表
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def partition(self, head: ListNode, x: int) -> ListNode:
        small, bigHead = ListNode(-1000), ListNode(-1000)
        res, big = small, bigHead
        while head != None:
            if head.val < x:
                small.next = head
                small = small.next
            else:
                bigHead.next = head
                bigHead = bigHead.next
            head = head.next
        small.next = big.next
        bigHead.next = None
        return res.next
