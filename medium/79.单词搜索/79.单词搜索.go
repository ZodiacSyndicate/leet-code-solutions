/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	s := board[i][j]
	board[i][j] = '*'
	res := dfs(board, i+1, j, k+1, word) ||
		dfs(board, i-1, j, k+1, word) ||
		dfs(board, i, j+1, k+1, word) ||
		dfs(board, i, j-1, k+1, word)
	board[i][j] = s
	return res
}
