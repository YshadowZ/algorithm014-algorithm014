package main

func main() {
	rotate([]int{1, 2, 3, 4}, 3)
}

// 暴力解法
// 两层循环
// 一层： 需要旋转的次数
// 二层： 每个元素移位
// 时间复杂度： O（n平方） 空间复杂度： O(1)
func rotate(nums []int, k int) {
	length := len(nums) // 数组总长度
	for moveCount := 0; moveCount < k; moveCount++ {
		num := nums[length-1] // 每次尾部都是要转移的那个数
		// 数组中的每个元素依次往后移位
		for index := length - 1; index > -1; index-- {
			// 第0个元素为尾部要转移的那个数
			if index > 0 {
				nums[index] = nums[index-1]
			} else {
				nums[index] = num
			}
		}
	}
}
