package lc56

import (
	"sort"
)

// Merge .
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	var bottom, top int = intervals[0][0], intervals[0][1]
	for index := range intervals {
		entity := &intervals[index]
		if (*entity)[0] <= top {
			if (*entity)[1] >= top {
				top = (*entity)[1]
			}
			continue
		}
		result = append(result, []int{bottom, top})
		bottom, top = (*entity)[0], (*entity)[1]
	}
	result = append(result, []int{bottom, top})
	return result
}
