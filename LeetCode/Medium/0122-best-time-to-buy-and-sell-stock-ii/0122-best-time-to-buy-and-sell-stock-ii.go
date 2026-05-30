func maxProfit(prices []int) int {
    profit := 0
    n := len(prices)

    for i := 0; i < n - 1; i++ {
        if (prices[i+1] - prices[i] > 0) {
            profit += prices[i+1] - prices[i]
        }
    }
    return profit
}