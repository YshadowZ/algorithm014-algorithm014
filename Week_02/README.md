学习笔记

# hashMap总结 数据结构即算法

## 介绍
hashMap表示的是一种键值对应关系，由哈希公式(用来映射键和存储数组位置的索引)和底层存储结构组成，其底层存储结构为数组和链表。和数组的区别在于：
* 数组的键是从零开始的连续的自然数，而hashMap的键可以是包括数字、字符串、布尔值、结构体多种数据类型（但都要通过一个哈希函数换算为一个数组索引）。
* 数组的查找复杂度为O(1), 插入、删除的复杂度为O(n), hashMap的查找、插入、删除复杂度正常情况下都为O(1)，极端查询复杂度是O(n),此时已退化为链表，因为是根据哈希函数计算后直接找到索引值。

## 关键点
* 哈希函数
* 冲突解决方法

哈希函数的运算性能决定了hashMap的读写性能，理想情况下，哈希函数应该能够将不同的键映均匀的射为底层存储数组中不同的索引位置，但由于底层数组不可能无限大， 而键值可以是无限多的，所以必然会出现映射后索引重复的问题，这时就是用一个链表来存储后续的值。

对于使用key查询的场景, 应该优先选择hashMap，数据结构即算法。


# 树

## 数据结构

```golang
type treeNode struct {
	val       int
	leftNode  *treeNode
	rightNode *treeNode
}
```

## 遍历

* 前序遍历 根-左-右
* 中序遍历 左-根-右
* 后序遍历 左-右-根

树本身就是具有重复性结构的数据结构，便于使用递归的方式遍历所有节点的值

```golang
// Preorder 树 前序遍历 根-左-右
func Preorder(root *treeNode) {
	if root != nil {
		fmt.Println(root.val)
		Preorder(root.leftNode)
		Preorder(root.rightNode)
	}
}
```

```golang
// Middleorder 中序遍历 左-根-右
func Middleorder(root *treeNode) {
	if root != nil {
		Middleorder(root.leftNode)
		fmt.Println(root.val)
		Middleorder(root.rightNode)
	}
}
```

```golang
// Postorder 后序遍历 左-右-根
func Postorder(root *treeNode) {
	if root != nil {
		Postorder(root.leftNode)
		Postorder(root.rightNode)
		fmt.Println(root.val)
	}
}
```

### 二叉树
每个根节点都有两个儿子节点

### 二叉搜索树
特点：排序过的树
* 左子树上**所有结点**的值均小于它的根节点的值
* 右子树上**所有节点**的值均大于它的根节点的值
* 中序遍历为升序遍历
* 插入和搜索都为O(log(n))的时间复杂度，但特殊情况下，二叉搜索树会退化成单链表(比如一致插入比根节点大的数值)，插入和搜索的时间复杂度会退化为O(n)
