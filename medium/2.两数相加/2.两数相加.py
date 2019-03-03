#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#
# https://leetcode-cn.com/problems/add-two-numbers/description/
#
# algorithms
# Medium (32.31%)
# Total Accepted:    87.9K
# Total Submissions: 268.4K
# Testcase Example:  '[2,4,3]\n[5,6,4]'
#
# 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
#
# 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
#
# 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
#
# 示例：
#
# 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
# 输出：7 -> 0 -> 8
# 原因：342 + 465 = 807
#
#
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = ListNode(0)
        node = head
        carry = False
        while l1 and l2:
            val = 0
            if carry:
                val = l1.val + l2.val + 1
                if val < 10:
                    carry = False
            else:
                val = l1.val + l2.val
                if val > 9:
                    carry = True
            val %= 10
            node.next = ListNode(val)
            node = node.next
            l1 = l1.next
            l2 = l2.next
        if l1:
            while l1:
                if carry:
                    if l1.val + 1 < 10:
                        carry = False
                    node.next = ListNode((l1.val + 1) % 10)
                    node = node.next
                    l1 = l1.next
                else:
                    node.next = l1
                    break
        if l2:
            while l2:
                if carry:
                    if l2.val + 1 < 10:
                        carry = False
                    node.next = ListNode((l2.val + 1) % 10)
                    node = node.next
                    l2 = l2.next
                else:
                    node.next = l2
                    break
        if carry and l1 == None and l2 == None:
            node.next = ListNode(1)
        return head.next
