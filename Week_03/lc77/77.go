package lc77

// https://leetcode-cn.com/problems/combinations/

// Combine 查找前k个数的组合
func Combine(n int, k int) [][]int {
	var arr []int
	var result [][]int
	if n == 0 || k == 0 {
		return result
	}
	for index := 1; index <= k; index++ {
		arr = append(arr, index)
	}
	getResult(n, k, 1, arr, &result)
	return result
}

// 不断从前k，k-1.。。。一直到前1个数为终止条件
func getResult(n, k, start int, arr []int, result *[][]int) {
	l := len(arr)
	for i := start; i <= n; i++ {
		if k == 1 {
			arr[l-1] = i
			dst := make([]int, l)
			copy(dst, arr)
			*result = append(*result, dst)
		} else {
			arr[l-k] = i
			getResult(n, k-1, i+1, arr, result)
		}
	}
}
