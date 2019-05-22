/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (57.84%)
 * Likes:    158
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 22.1K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(pl, pr, il, ir int) *TreeNode
	build = func(pl, pr, il, ir int) *TreeNode {
		if pl > pr || il > ir {
			return nil
		}
		node := new(TreeNode)
		node.Val = preorder[pl]
		index := 0
		for i := il; i <= ir; i++ {
			if inorder[i] == preorder[pl] {
				index = i
				break
			}
		}
		lCount := index - il
		node.Left = build(pl+1, pl+lCount, il, index-1)
		node.Right = build(pl+lCount+1, pr, index+1, ir)
		return node
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

