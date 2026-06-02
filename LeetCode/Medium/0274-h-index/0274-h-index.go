func hIndex(citations []int) int {
    idx := 0
    
    n := len(citations)
    slices.Sort(citations)

    for i := n - 1; i >= 0; i-- {
        h := min(n - i, citations[i])
        idx = max(idx, h)
    }

    return idx
}