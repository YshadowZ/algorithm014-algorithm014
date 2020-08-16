// https://leetcode-cn.com/problems/trapping-rain-water/
package lc42

// 从底向上走，实心为1，空心为0，被1夹着的0即为可用容积
// 时间复杂度O(n平方) 空间复杂度O(1)

func trap(height []int) int {
	length := len(height)
	var max, area int
	for index := 0; index < length; index++ {
		if height[index] > max {
			max = height[index]
		}
	}
	for out := 1; out <= max; out++ {
		left := false
		validArea := 0
		for inner := 0; inner < length; inner++ {
			if height[inner] < out {
				if left {
					validArea++
				}
			} else {
				if validArea > 0 && left {
					area += validArea
					validArea = 0
				}
				left = true
			}
		}
	}
	return area
}
