/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (60.75%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    7.9K
 * Total Submissions: 13K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var build func(il, ir, pl, pr int) *TreeNode
	build = func(il, ir, pl, pr int) *TreeNode {
		if il > ir || pl > pr {
			return nil
		}
		node := new(TreeNode)
		node.Val = postorder[pr]
		index := 0
		for i := il; i <= ir; i++ {
			if postorder[pr] == inorder[i] {
				index = i
				break
			}
		}
		rCount := ir - index
		node.Left = build(il, index-1, pl, pr-rCount-1)
		node.Right = build(index+1, ir, pr-rCount, pr-1)
		return node
	}
	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

