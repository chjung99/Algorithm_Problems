func sortTheStudents(score [][]int, k int) [][]int {
    slices.SortFunc(score, func(a, b []int) int {
        return cmp.Compare(b[k], a[k])
    })
    return score
}