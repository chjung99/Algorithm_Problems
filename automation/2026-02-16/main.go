func titleToNumber(columnTitle string) int {
    s := 0
    n := len(columnTitle)

    m := make(map[byte]int)
    for i := 0; i < 26; i++ {
        m[byte('A'+i)] = i+1
    }

    fac := 1
    const NUM_OF_ALPHABET = 26
    for i := n-1; i >= 0; i-- {
        val := m[columnTitle[i]]
        s += fac * val
        fac *= NUM_OF_ALPHABET
    }
    return s

}