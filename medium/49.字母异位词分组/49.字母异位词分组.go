/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (54.79%)
 * Total Accepted:    13.6K
 * Total Submissions: 24.7K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	ans := map[[]int] []string{}
	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			count[c - 'a']++
		}
		if _, ok := ans[count]; ok {
			ans[count] = append(ans[count], s)
		} else {
			ans[count] = []string{s}
		}
	}
	result := [][]string{}
	for _, item := range ans {
		result = append(result, item)
	}
	return result
}

