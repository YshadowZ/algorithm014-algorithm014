package lc45

// https://leetcode-cn.com/problems/jump-game-ii/

// 需要优化，内存占用过高
func Jump(nums []int) int {
	type node struct {
		index int
		level int
	}
	start := &node{index: 0, level: 0}
	end := &node{index: len(nums) - 1}
	queue := []*node{start}
	deep := 0
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for i := 1; i <= nums[current.index]; i++ {
			if current.index+i >= len(nums) {
				break
			}
			child := &node{index: current.index + i, level: current.level + 1}
			if child.index == end.index {
				deep = child.level
				break
			}
			queue = append(queue, child)
		}
		if deep > 0 {
			break
		}
	}
	return deep
}
