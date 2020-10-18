package main

import (
	// "week_08/lc146"
	"fmt"
	"week_09/lc300"
)

func main() {
	// lru := lc146.Constructor(10)
	// lru.Put(10, 13)
	// lru.Put(3, 17)
	// lru.Put(6, 11)
	// lru.Put(10, 5)
	// lru.Get(1)
	result := lc300.LengthOfLIS([]int{2, 3, 4, 10, 1, 2, 3, 4})
	fmt.Println(result)
}
