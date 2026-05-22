func rob(nums []int) int {
    n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	if n >= 2 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}