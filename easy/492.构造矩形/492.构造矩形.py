#
# @lc app=leetcode.cn id=492 lang=python3
#
# [492] 构造矩形
#
# https://leetcode-cn.com/problems/construct-the-rectangle/description/
#
# algorithms
# Easy (44.54%)
# Total Accepted:    2.3K
# Total Submissions: 5.2K
# Testcase Example:  '1'
#
# 作为一位web开发者， 懂得怎样去规划一个页面的尺寸是很重要的。 现给定一个具体的矩形页面面积，你的任务是设计一个长度为 L 和宽度为 W
# 且满足以下要求的矩形的页面。要求：
#
#
# 1. 你设计的矩形页面必须等于给定的目标面积。
#
# 2. 宽度 W 不应大于长度 L，换言之，要求 L >= W 。
#
# 3. 长度 L 和宽度 W 之间的差距应当尽可能小。
#
#
# 你需要按顺序输出你设计的页面的长度 L 和宽度 W。
#
# 示例：
#
#
# 输入: 4
# 输出: [2, 2]
# 解释: 目标面积是 4， 所有可能的构造方案有 [1,4], [2,2], [4,1]。
# 但是根据要求2，[1,4] 不符合要求; 根据要求3，[2,2] 比 [4,1] 更能符合要求. 所以输出长度 L 为 2， 宽度 W 为 2。
#
#
# 说明:
#
#
# 给定的面积不大于 10,000,000 且为正整数。
# 你设计的页面的长度和宽度必须都是正整数。
#
#
#
import math


class Solution:
    def constructRectangle(self, area: int) -> List[int]:
        mid = math.ceil(math.sqrt(area))
        res = []
        while mid > 0:
            if area % mid == 0 and area // mid >= mid:
                res = [area // mid, mid]
                break
            mid -= 1
        return res
