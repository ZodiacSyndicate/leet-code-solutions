/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (63.18%)
 * Total Accepted:    13.7K
 * Total Submissions: 21.6K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	result := [][]int{}
	arr := []int{}
	var combination func(temp []int, target int, k int)
	combination = func(temp []int, target int, k int) {
		if target == 0 {
			result = append(result, temp)
			return
		}
		if target < candidates[0] {
			return
		}
		for i := k; i < len(candidates) && candidates[i] <= target; i++ {
			list := make([]int, len(temp))
			copy(list, temp)
			list = append(list, candidates[i])
			combination(list, target-candidates[i], i)
		}
	}
	combination(arr, target, 0)
	return result
}

