func digitFrequencyScore(n int) int {
    feq := make(map[int]int)
    score := 0

    for (n != 0) {
        x := n % 10
        n = n / 10

        _, exists := feq[x]
        if (!exists) {
            feq[x] = 0
        }
        feq[x] += 1
    }

    for key, val := range feq {
        score += key * val
    }
    return score
}