/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (36.80%)
 * Total Accepted:    10.7K
 * Total Submissions: 29.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
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
import "math"

type Item struct {
	Cur  float64
	Node *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := float64(^uint(0))
	item := new(Item)
	item.Cur = 1
	item.Node = root
	stack := []*Item{item}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Node != nil {
			left, right := new(Item), new(Item)
			left.Cur, right.Cur = node.Cur+1, node.Cur+1
			left.Node, right.Node = node.Node.Left, node.Node.Right
			stack = append(stack, left)
			stack = append(stack, right)
			if node.Node.Left == nil && node.Node.Right == nil {
				res = math.Min(node.Cur, res)
			}
		}
	}
	return int(res)
}
