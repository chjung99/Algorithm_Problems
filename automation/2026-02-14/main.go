func isIsomorphic(s string, t string) bool {
    n := len(s)
    m := make(map[byte]byte)
    u := make(map[byte]int)
    x := []byte(t)

    for i := 0; i < n; i++ {
        _, exist := m[t[i]]
        if (!exist) {
            _, isUsed := u[s[i]]
            if (isUsed) {
                return false
            }
            m[t[i]] = s[i]
            u[s[i]] = 1
        }

        x[i] = m[t[i]]
    }
    return s == string(x)
}