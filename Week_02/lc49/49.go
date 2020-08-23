package lc49

// 字母异位词分组
// https://leetcode-cn.com/problems/group-anagrams/

import "sort"

// GroupAnagrams .
// 每一个元素按照字符排序，看排序后的值是否再一个hashmap中存在
// 不存在，作为一个数组存入，存在则append到数组中
// 最后便利hashMap, 得到结果
// 时间复杂度O(K * nlog(n)), 空间复杂度O(n)
func GroupAnagrams(strs []string) [][]string {
	stringMap := map[string][]string{}
	result := [][]string{}
	for _, s := range strs {
		cArray := []rune(s)
		sort.SliceStable(cArray, func(i, j int) bool {
			return cArray[i] < cArray[j]
		})
		if len(stringMap[string(cArray)]) != 0 {
			stringMap[string(cArray)] = append(stringMap[string(cArray)], s)
		} else {
			stringMap[string(cArray)] = []string{s}
		}
	}
	for _, val := range stringMap {
		result = append(result, val)
	}
	return result
}
