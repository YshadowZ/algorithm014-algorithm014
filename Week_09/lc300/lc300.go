package lc300

func LengthOfLIS(nums []int) int {
	targetLength := len(nums)
	if targetLength == 0 {
		return 0
	}
	dp := make([]int, targetLength)
	dp[0] = 1
	var maxResult int = 1
	for i := 1; i < targetLength; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxResult = max(maxResult, dp[i])
	}
	return maxResult
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
