#
# @lc app=leetcode.cn id=92 lang=python3
#
# [92] 反转链表 II
#
# https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
#
# algorithms
# Medium (42.49%)
# Likes:    137
# Dislikes: 0
# Total Accepted:    10.4K
# Total Submissions: 23.8K
# Testcase Example:  '[1,2,3,4,5]\n2\n4'
#
# 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
#
# 说明:
# 1 ≤ m ≤ n ≤ 链表长度。
#
# 示例:
#
# 输入: 1->2->3->4->5->NULL, m = 2, n = 4
# 输出: 1->4->3->2->5->NULL
#
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        dummy = ListNode(-10000)
        dummy.next = head
        i = 1
        node = dummy
        while i < m:
            node = node.next
            i += 1
        p, q = None, None
        while i <= n:
            temp = node.next
            if q == None:
                q = temp
            node.next = node.next.next
            temp.next = p
            p = temp
            i += 1
        q.next = node.next
        node.next = p
        return dummy.next
