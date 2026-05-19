func maxProfit(prices []int) int {
    min := prices[0]
    profit := 0

    for i := 0; i < len(prices); i++ {
        diff := prices[i] - min 
        if (diff > profit) {
            profit = diff
        }
        if (prices[i] < min) {
            min = prices[i]
        }
    }

    return profit
}