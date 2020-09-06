package lc860

// https://leetcode-cn.com/problems/lemonade-change/description/

// 贪心算法，剩余价格一次使用10元，5元递减
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	if bills[0] > 5 {
		return false
	}
	smallChange5 := []int{}
	smallChange10 := []int{}
	for _, code := range bills {
		if code == 5 {
			smallChange5 = append(smallChange5, code)
			continue
		}
		result := giveChange(code-5, &smallChange5, &smallChange10)
		if result == false {
			return false
		}
		if code == 10 {
			smallChange10 = append(smallChange10, code)
		}
	}
	return true
}

func giveChange(remianCount int, smallChange5 *[]int, smallChange10 *[]int) bool {
	if remianCount > 5 {
		if len(*smallChange10) > 0 {
			*smallChange10 = (*smallChange10)[1:]
			return giveChange(remianCount-10, smallChange5, smallChange10)
		}
		if len(*smallChange5) > 0 {
			*smallChange5 = (*smallChange5)[1:]
			return giveChange(remianCount-5, smallChange5, smallChange10)
		}
		return false
	}
	if len(*smallChange5) > 0 {
		*smallChange5 = (*smallChange5)[1:]
		return true
	}
	return false
}
