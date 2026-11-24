func largestAltitude(gain []int) int {
    m := 0
    x := 0
    for _, num := range gain {
        x += num
        m = max(m, x)
    }
    return m
}