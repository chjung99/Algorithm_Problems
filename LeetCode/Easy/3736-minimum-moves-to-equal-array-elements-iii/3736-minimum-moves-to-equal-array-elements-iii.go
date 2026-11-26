func minMoves(nums []int) int {
    e := 0
    m := 0

    for _, num := range nums {
        e = max(e, num)
    }

    for _, num := range nums {
        m += e - num
    }
    return m
}