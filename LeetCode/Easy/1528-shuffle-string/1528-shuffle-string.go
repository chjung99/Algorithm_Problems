func restoreString(s string, indices []int) string {
    ret := ""
    n := len(s)

    m := make(map[int]byte)

    for i, idx := range indices {
        m[idx] = s[i]
    }

    for i := 0; i < n; i++ {
        ret += string(m[i])
    }

    return ret
}