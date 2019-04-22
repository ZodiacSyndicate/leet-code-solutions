/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (35.39%)
 * Total Accepted:    13K
 * Total Submissions: 36.8K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */
/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"

type Intervals []Interval

func (this Intervals) Len() int {
	return len(this)
}

func (this Intervals) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Intervals) Less(i, j int) bool {
	return this[i].Start < this[j].Start
}

func merge(intervals []Interval) []Interval {
	res := []Interval{}
	if len(intervals) == 0 {
		return res
	}
	sort.Sort(Intervals(intervals))
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if cur.End >= intervals[i].Start {
			cur.End = max(intervals[i].End, cur.End)
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
