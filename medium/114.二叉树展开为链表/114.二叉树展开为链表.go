/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (60.73%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    6.8K
 * Total Submissions: 11.2K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为链表。
 *
 * 例如，给定二叉树
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 * 将其展开为：
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	if root.Left == nil {
		return
	}
	r := root.Right
	l := root.Left
	node := l
	for node != nil && node.Right != nil {
		node = node.Right
	}
	node.Right = r
	root.Right = l
	root.Left = nil
}

