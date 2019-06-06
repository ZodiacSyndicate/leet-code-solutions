/*
 * @lc app=leetcode.cn id=127 lang=javascript
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
/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function(beginWord, endWord, wordList) {
  const dict = new Set(wordList)
  if (!dict.has(endWord)) return 0
  let [q1, q2] = [new Set([beginWord]), new Set([endWord])]
  let l = beginWord.length
  let step = 0
  while (q1.size && q2.size) {
    step++
    if (q1.size > q2.size) {
      ;[q1, q2] = [q2, q1]
    }
    let q = new Set()
    for (let w of q1.values()) {
      for (let i = 0; i < l; i++) {
        const c = w[i]
        for (let j = 'a'.charCodeAt(); j < 'z'.charCodeAt(); j++) {
          w = w.slice(0, i) + String.fromCharCode(j) + w.slice(i + 1)
          if (q2.has(w)) {
            return step + 1
          }
          if (!dict.has(w)) continue
          dict.delete(w)
          q.add(w)
        }
        w = w.slice(0, i) + c + w.slice(i + 1)
      }
    }
    ;[q, q1] = [q1, q]
  }
  return 0
}
