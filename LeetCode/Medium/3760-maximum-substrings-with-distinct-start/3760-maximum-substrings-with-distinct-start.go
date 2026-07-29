func maxDistinct(s string) int {
    m := make(map[rune]int)

    for _, c := range s {
        _, exists := m[c]
        if (!exists) {
            m[c] = 0
        }
        m[c] += 1
    }
    return len(m)
}