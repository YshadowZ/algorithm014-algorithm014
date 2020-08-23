package tree

import "fmt"

type treeNode struct {
	val       int
	leftNode  *treeNode
	rightNode *treeNode
}

// Preorder 树 前序遍历 根-左-右
func Preorder(root *treeNode) {
	if root != nil {
		fmt.Println(root.val)
		Preorder(root.leftNode)
		Preorder(root.rightNode)
	}
}

// Middleorder 中序遍历 左-根-右
func Middleorder(root *treeNode) {
	if root != nil {
		Middleorder(root.leftNode)
		fmt.Println(root.val)
		Middleorder(root.rightNode)
	}
}

// Postorder 后序遍历 左-右-根
func Postorder(root *treeNode) {
	if root != nil {
		Postorder(root.leftNode)
		Postorder(root.rightNode)
		fmt.Println(root.val)
	}
}
