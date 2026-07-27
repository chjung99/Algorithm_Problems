func mirrorDistance(n int) int {
    rev := 0
    x := n

    for x != 0 {
        rev *= 10
        rev += x % 10
        x = x / 10
    }

    return max(n - rev, rev - n)
}