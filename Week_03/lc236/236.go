package lc236

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 1.如果节点的值和左或右的节点值相同，那么这个节点就是p或q的祖先
// 2.否则就继续顺着左子树或者右子树继续向下找
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftNode := LowestCommonAncestor(root.Left, p, q)
	rightNode := LowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode != nil {
		return leftNode
	}
	if rightNode != nil {
		return rightNode
	}
	return nil
}
