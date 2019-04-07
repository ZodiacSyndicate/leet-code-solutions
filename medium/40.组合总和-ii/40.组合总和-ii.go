/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (54.32%)
 * Total Accepted:    10.4K
 * Total Submissions: 19.1K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的每个数字在每个组合中只能使用一次。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 */
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
  ans, arr := [][]int{}, []int{}
  sort.Sort(sort.IntSlice(candidates))

  var combination func(temp []int, target int, k int)
  combination = func (temp []int, target int, k int) {
    if target == 0 {
      ans = append(ans, temp)
      return
    }
    i := k
    for i < len(candidates) && candidates[i] <= target {
      list := temp[:]
      list = append(list, candidates[i])
      combination(list, target - candidates[i], i + 1)
      for i < len(candidates) - 1 && candidates[i + 1] == candidates[i] {
        i++
      }
      i++
    }
  }
  combination(arr, target, 0)
  return ans
}

