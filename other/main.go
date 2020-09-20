package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	// lru := Constructor(10)
	// lru.Put(10, 13)
	// lru.Put(3, 17)
	// lru.Put(6, 11)
	// lru.Put(10, 5)
	// lru.Get(1)

	// lru.Put(3, 3)
	// lru.Get(2)
	// lru.Put(4, 4)
	// lru.Get(1)
	// lru.Get(3)
	// lru.Get(4)
	// lru.Put(1, 20)
	// lru.Get(1)
	// lru.Put(10, 10)
	// lru.Put(11, 11)
	// lru.Get(1)
}

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
			// fmt.Println(l.RecordMap[key].Value)
			l.ShowALl()
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
		// fmt.Println(l.RecordMap[key].Value)
		l.ShowALl()
		return l.RecordMap[key].Value
	}
	// fmt.Println(-1)
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
	l.ShowALl()
	// fmt.Println(l)
}

func (l *LRUCache) ShowALl() {
	node := l.Head
	for node != nil {
		fmt.Print(node.Value, ": ", node, ", ")
		node = node.Next
	}
	fmt.Print("\n")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func moveZeroes(nums []int) {
	length := len(nums)
	for i, j := 0, length-1; i < j; i++ {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				for nums[j] == 0 {
					j--
				}
				if j <= i {
					break
				} else {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
	fmt.Println(nums)
}
