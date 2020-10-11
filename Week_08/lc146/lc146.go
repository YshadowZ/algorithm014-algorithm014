package lc146

// LRUCache .
type LRUCache struct {
	RecordMap map[int]*LListNode
	Head      *LListNode
	Tail      *LListNode
	Length    int
	Cap       int
}

// LListNode .
type LListNode struct {
	Prev  *LListNode
	Value int
	Key   int
	Next  *LListNode
}

// Constructor 构造
func Constructor(capacity int) LRUCache {
	result := LRUCache{Cap: capacity, RecordMap: map[int]*LListNode{}, Head: nil, Tail: nil, Length: 0}
	return result
}

// Get 获取
func (l *LRUCache) Get(key int) int {
	if l.Cap == 0 {
		return -1
	}
	if l.RecordMap[key] != nil {
		if l.Head == l.RecordMap[key] {
			return l.RecordMap[key].Value
		}
		if l.Tail == l.RecordMap[key] {
			l.Tail = l.RecordMap[key].Prev
		}
		if l.RecordMap[key].Prev != nil {
			l.RecordMap[key].Prev.Next = l.RecordMap[key].Next
		}
		if l.RecordMap[key].Next != nil {
			l.RecordMap[key].Next.Prev = l.RecordMap[key].Prev
		}
		l.RecordMap[key].Next = l.Head
		l.RecordMap[key].Prev = nil
		l.Head.Prev = l.RecordMap[key]
		l.Head = l.RecordMap[key]
		return l.RecordMap[key].Value
	}
	return -1
}

// Put 更新
func (l *LRUCache) Put(key int, value int) {
	if l.Cap == 0 {
		return
	}
	node := LListNode{Prev: nil, Value: value, Key: key, Next: nil}
	if l.RecordMap[key] == nil {
		if l.Head == nil {
			l.Head = &node
			l.Tail = &node
		} else {
			node.Next = l.Head
			l.Head.Prev = &node
			l.Head = &node
		}
		l.Length++
		l.RecordMap[key] = &node
		if l.Length > l.Cap {
			delete(l.RecordMap, l.Tail.Key)
			l.Tail.Prev.Next = nil
			l.Tail = l.Tail.Prev
			l.Length--
		}
	} else {
		l.RecordMap[key].Value = value
		if l.Tail == l.RecordMap[key] {
			l.Tail = l.RecordMap[key].Prev
		}
		if l.RecordMap[key].Prev != nil {
			l.RecordMap[key].Prev.Next = l.RecordMap[key].Next
		}
		if l.RecordMap[key].Next != nil {
			l.RecordMap[key].Next.Prev = l.RecordMap[key].Prev
		}
		l.RecordMap[key].Next = l.Head
		l.RecordMap[key].Prev = nil
		l.Head.Prev = l.RecordMap[key]
		l.Head = l.RecordMap[key]
	}
}

// func (l *LRUCache) ShowALl() {
// 	node := l.Head
// 	for node != nil {
// 		fmt.Print(node.Value, ": ", node, ", ")
// 		node = node.Next
// 	}
// 	fmt.Print("\n")
// }
