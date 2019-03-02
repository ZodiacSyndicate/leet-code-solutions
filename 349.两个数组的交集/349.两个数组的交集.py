#
# @lc app=leetcode.cn id=349 lang=python3
#
# [349] 两个数组的交集
#
# https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
#
# algorithms
# Easy (59.17%)
# Total Accepted:    13.6K
# Total Submissions: 22.8K
# Testcase Example:  '[1,2,2,1]\n[2,2]'
#
# 给定两个数组，编写一个函数来计算它们的交集。
#
# 示例 1:
#
# 输入: nums1 = [1,2,2,1], nums2 = [2,2]
# 输出: [2]
#
#
# 示例 2:
#
# 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
# 输出: [9,4]
#
# 说明:
#
#
# 输出结果中的每个元素一定是唯一的。
# 我们可以不考虑输出结果的顺序。
#
#
#


class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        m = {}
        for n in nums1:
            if not n in m:
                m[n] = 1
        res = []
        for n in nums2:
            if n in m:
                res.append(n)
                del m[n]
        return res
