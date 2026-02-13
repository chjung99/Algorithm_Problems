func isIsomorphic(s string, t string) bool {
    n := len(s)
    m := make(map[byte]byte)
    x := []byte(t)

    for i := 0; i < n; i++ {
        _, exist := m[s[i]]
        if (!exist) {
            m[s[i]] = t[i]
        }
        x[i] = m[s[i]]
    }
    return t == string(x)
}