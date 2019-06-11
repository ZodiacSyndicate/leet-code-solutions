/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (34.02%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    5.8K
 * Total Submissions: 17K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := map[string]string{}
	for _, s := range wordList {
		dict[s] = s
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	q1 := map[string]string{beginWord: beginWord}
	q2 := map[string]string{endWord: endWord}
	l := len(beginWord)
	step := 0
	for len(q1) > 0 && len(q2) > 0 {
		step++
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		q := map[string]string{}
		for w := range q1 {
			for i := 0; i < l; i++ {
				c := string(w[i])
				for j := 'a'; j < 'z'; j++ {
					w = w[0:i] + string(j) + w[i+1:]
					if _, ok := q2[w]; ok {
						return step + 1
					}
					if _, ok := dict[w]; !ok {
						continue
					}
					delete(dict, w)
					q[w] = w
				}
				w = w[0:i] + c + w[i+1:]
			}
		}
		q, q1 = q1, q
	}
	return 0
}

