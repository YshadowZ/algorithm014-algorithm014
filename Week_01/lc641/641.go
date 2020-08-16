// https://leetcode-cn.com/problems/design-circular-deque/

package lc641

// MyCircularDeque 双端队列
type MyCircularDeque struct {
	Limit int
	First *Node
	Last  *Node
	Sum   int
}

// Node 节点
type Node struct {
	Value int
	Next  *Node
	Pre   *Node
}

// Constructor Initialize your data structure here. Set the size of the deque to be k.
// Constructor 构造函数
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{Limit: k, Sum: 0}
}

// InsertFront /** Adds an item at the front of Deque. Return true if the operation is successful. */
// InsertFront 头部插入一个元素
func (m *MyCircularDeque) InsertFront(value int) bool {
	if m.IsFull() {
		return false
	}
	node := Node{Value: value}
	if m.First == nil {
		m.First = &node
		m.Last = &node
	} else {
		node.Next = m.First
		m.First.Pre = &node
		m.First = &node
	}
	m.Sum++
	return true
}

// InsertLast /** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) InsertLast(value int) bool {
	if m.IsFull() {
		return false
	}
	node := Node{Value: value}
	if m.Last == nil {
		m.First = &node
		m.Last = &node
	} else {
		m.Last.Next = &node
		node.Pre = m.Last
		m.Last = &node
	}
	m.Sum++
	return true
}

// DeleteFront /** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) DeleteFront() bool {
	if m.IsEmpty() {
		return false
	}
	if m.Sum == 1 {
		m.First = nil
		m.Last = nil
	} else {
		m.First = m.First.Next
		m.First.Pre = nil
	}
	m.Sum--
	return true
}

// DeleteLast /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) DeleteLast() bool {
	if m.IsEmpty() {
		return false
	}
	if m.Sum == 1 {
		m.First = nil
		m.Last = nil
	} else {
		m.Last = m.Last.Pre
		m.Last.Next = nil
	}
	m.Sum--
	return true
}

// GetFront /** Get the front item from the deque. */
func (m *MyCircularDeque) GetFront() int {
	if m.IsEmpty() {
		return -1
	}
	result := m.First.Value
	// m.DeleteFront()
	return result
}

// GetRear /** Get the last item from the deque. */
func (m *MyCircularDeque) GetRear() int {
	if m.IsEmpty() {
		return -1
	}
	result := m.Last.Value
	// m.DeleteLast()
	return result
}

// IsEmpty /** Checks whether the circular deque is empty or not. */
func (m *MyCircularDeque) IsEmpty() bool {
	return m.Sum == 0
}

// IsFull /** Checks whether the circular deque is full or not. */
func (m *MyCircularDeque) IsFull() bool {
	return m.Sum == m.Limit
}
