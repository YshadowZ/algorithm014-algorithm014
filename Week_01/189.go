package main

func main() {
	// rotate([]int{1, 2, 3, 4}, 3)
	rotate_2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
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

// 优化版本
// 移动次数对数组长度取余，为0的情况下不需要任何行动
// 反转整个数组
// 反转前k个数字
// 反转后面的数字
// 时间、空间复杂度都为O(1)
// 比较难想到的算法，多在纸上画画应该会更好找规律
func rotate_2(nums []int, k int) {
	length := len(nums) // 数组总长度
	k = k % length
	if k == 0 {
		return
	}
	reverse(&nums, 0, length-1)
	reverse(&nums, 0, k-1)
	reverse(&nums, k, length-1)
}

func reverse(nums *[]int, start, end int) {
	length := end - start + 1
	loopTime := 0
	if length%2 == 1 {
		loopTime = (length - 1) / 2
	} else {
		loopTime = length / 2
	}
	for index := 0; index < loopTime; index++ {
		(*nums)[start+index], (*nums)[length-index-1+start] = (*nums)[length-index-1+start], (*nums)[start+index]
	}
}
