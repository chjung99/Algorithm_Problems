func isSubsequence(s string, t string) bool {
    sPtr := 0
    tPtr := 0
    sLen := len(s)
    tLen := len(t)

    for (tPtr != tLen && sPtr != sLen) {
        if (s[sPtr] == t[tPtr]) {
            sPtr += 1
        }
        tPtr += 1
    }
    return sPtr == sLen
}
