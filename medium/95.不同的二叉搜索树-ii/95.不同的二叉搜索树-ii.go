/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (54.81%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 10.1K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		subLeft := generate(start, i-1)
		subRight := generate(i+1, end)
		for _, left := range subLeft {
			for _, right := range subRight {
				node := new(TreeNode)
				node.Val = i
				node.Left = left
				node.Right = right
				res = append(res, node)
			}
		}
	}
	return res
}

