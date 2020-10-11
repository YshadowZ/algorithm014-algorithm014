package main

import (
	// "week_08/lc146"
	"fmt"
	"week_08/lc56"
)

func main() {
	// lru := lc146.Constructor(10)
	// lru.Put(10, 13)
	// lru.Put(3, 17)
	// lru.Put(6, 11)
	// lru.Put(10, 5)
	// lru.Get(1)
	result := lc56.Merge([][]int{{1, 4}, {4, 5}})
	fmt.Println(result)
}
