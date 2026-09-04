func maximumWealth(accounts [][]int) int {
    maximumWealth := 0
    m := len(accounts)
    n := len(accounts[0])

    for i := range m {
        wealth := 0
        for j := range n {
            wealth += accounts[i][j]
        }
        maximumWealth = max(maximumWealth, wealth)
    }
    return maximumWealth
}