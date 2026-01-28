func strStr(haystack string, needle string) int {
    n := len(haystack)
    m := len(needle)

    if (n < m) {
        return -1
    }

    for i, _ := range(haystack) {
        if (i+m >= n) {
            break
        }
        if (haystack[i:i+m] == needle) {
            return i
        }
    }
    return -1
}