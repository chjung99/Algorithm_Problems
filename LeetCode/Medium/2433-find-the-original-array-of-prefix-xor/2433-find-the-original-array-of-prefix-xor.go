func findArray(pref []int) []int {
    arr := make([]int, 0)

    arr = append(arr, pref[0])
    for i := 1; i < len(pref); i++ {
        arr = append(arr, pref[i-1]^pref[i])
    }
    return arr
}