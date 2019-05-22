/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (55.05%)
 * Likes:    86
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 16.4K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 *
 *
 * 返回:
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
 *
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
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var dfs func(node *TreeNode, sum int, temp []int)
	dfs = func(node *TreeNode, sum int, temp []int) {
		arr := append([]int{}, temp...)
		arr = append(arr, node.Val)
		if node.Left == nil && node.Right == nil && sum-node.Val == 0 {
			res = append(res, arr)
			return
		}
		if node.Left != nil {
			dfs(node.Left, sum-node.Val, arr)
		}
		if node.Right != nil {
			dfs(node.Right, sum-node.Val, arr)
		}
	}
	dfs(root, sum, []int{})
	return res
}

