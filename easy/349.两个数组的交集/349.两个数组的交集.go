/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (59.17%)
 * Total Accepted:    13.6K
 * Total Submissions: 22.8K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, n := range nums1 {
		if _, ok := m[n]; ok == false {
			m[n] = 1
		}
	}
	res := []int{}
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			res = append(res, n)
			delete(m, n)
		}
	}
	return res
}

