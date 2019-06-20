/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (41.03%)
 * Likes:    119
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 24.9K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 *
 * 说明：
 *
 *
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 */
func wordBreak(s string, wordDict []string) bool {
	m := map[string]int{}
	memo := make([]int, len(s))
	for _, word := range wordDict {
		m[word] = 1
	}
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == len(s) {
			return true
		}
		if memo[start] != 0 {
			if memo[start] == 1 {
				return false
			}
			return true
		}
		for end := start + 1; end <= len(s); end++ {
			if _, ok := m[s[start:end]]; ok && dfs(end) {
				memo[start] = 2
				return true
			}
		}
		memo[start] = 1
		return false
	}
	return dfs(0)
}

