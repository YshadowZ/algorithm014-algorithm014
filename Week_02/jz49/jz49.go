// 丑数
// https://leetcode-cn.com/problems/chou-shu-lcof/
package jz49

import (
	"fmt"
	"sort"
	"time"
)

// 暴力求解
// 从1开始计算，计算每一个数的质因子是否只有2，3，5，如果是，则目标数就加一，一直求解到指定个数
// 空间复杂度O(1), 时间复杂度O(n平方)，求解个数多的话，leetcode会超时
func NthUglyNumber(n int) int {
	var index int = 0
	var num int = 0
	for i := 1; ; i++ {
		if i == 1 {
			index = 1
			num = 1
		} else {
			temp := i
			loop := true
			for loop {
				if temp%2 == 0 {
					temp = temp / 2
				} else if temp%3 == 0 {
					temp = temp / 3
				} else if temp%5 == 0 {
					temp = temp / 5
				} else {
					break
				}
			}
			if temp == 1 {
				index++
				num = i
			}
		}
		if index == n {
			break
		}
	}
	return num
}

// 优化方案
// 以1为基数，分别乘以2，3，5，得到[2,3,5],  选择最小的那个数为第二个数，同时移除[2,3,5]中的最小数
// 以第二个数为基数，分别乘以2,3,5,得到[4,6,10], 结合上一个迭代余下的数，选择一个最小数作为第三个数，同时移除所有和最小数相等的数
// 时间复杂度O(K * nlog(n)), 空间复杂度O(n)
// 输入为1690时，golang显示程序运行时间为33ms，但leetcode会显示超时
func NthUglyNumber_1(n int) int {
	if n == 1 {
		return 1
	}
	t1 := time.Now()
	var num int = 1
	temp := []int{}
	entities := []int{2, 3, 5}
	for i := 2; ; i++ {
		for _, a := range entities {
			temp = append(temp, num*a)
		}
		sort.Ints(temp)
		if i == n {
			totalTime := time.Since(t1)
			fmt.Println(totalTime)
			return temp[0]
		}
		num = temp[0]
		for i := 1; ; i++ {
			if temp[i] != num {
				temp = temp[i:]
				break
			}
		}
	}
}
