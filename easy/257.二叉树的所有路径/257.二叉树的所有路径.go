/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (56.91%)
 * Total Accepted:    5.9K
 * Total Submissions: 10.4K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
import "strconv"
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	return buildPath(res, root, strconv.Itoa(root.Val))
}

func buildPath(arr []string, node *TreeNode, cur string) []string {
	if node.Left == nil && node.Right == nil {
		arr = append(arr, cur)
	}
	if node.Left != nil {
		arr = buildPath(arr, node.Left, cur + "->" + strconv.Itoa(node.Left.Val))
	}
	if node.Right != nil {
		arr = buildPath(arr, node.Right, cur + "->" + strconv.Itoa(node.Right.Val))
	}
	return arr
}

