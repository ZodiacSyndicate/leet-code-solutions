/*
 * @lc app=leetcode.cn id=79 lang=javascript
 *
 * [79] 单词搜索
 */
/**
 * @param {character[][]} board
 * @param {string} word
 * @return {boolean}
 */
var exist = function(board, word) {
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      if (dfs(board, i, j, word, 0)) return true
    }
  }
  return false
}

/**
 *
 * @param {string[][]} board
 * @param {number} i
 * @param {number} j
 * @param {string} word
 * @param {number} k
 */
function dfs(board, i, j, word, k) {
  if (k === word.length) return true
  if (
    i < 0 ||
    i >= board.length ||
    j < 0 ||
    j >= board[0].length ||
    board[i][j] !== word[k]
  )
    return false
  const s = board[i][j]
  board[i][j] = '*'
  const res =
    dfs(board, i + 1, j, word, k + 1) ||
    dfs(board, i - 1, j, word, k + 1) ||
    dfs(board, i, j + 1, word, k + 1) ||
    dfs(board, i, j - 1, word, k + 1)
  board[i][j] = s
  return res
}
