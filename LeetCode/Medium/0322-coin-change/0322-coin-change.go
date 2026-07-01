func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

    dp[0] = 1

    for i := 1; i <= amount; i++ {
        minCnt := -1
        for _, num := range coins {
            if (i - num < 0) {
                continue
            } else if (dp[i - num] == 0) {
                continue
            } else {
                if (minCnt == -1) {
                    minCnt = dp[i - num]
                } else {
                    minCnt = min(minCnt, dp[i-num])
                }
            }
        }
        // fmt.Println(i, minCnt)
        dp[i] = minCnt + 1
    }

    if (dp[amount] == 0) {
        return -1
    }
    return dp[amount] - 1
}