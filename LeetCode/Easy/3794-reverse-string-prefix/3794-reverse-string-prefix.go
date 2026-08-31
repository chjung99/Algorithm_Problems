func reversePrefix(s string, k int) string {
    rev := make([]byte, 0)

    for i := k-1; i >=0 ; i-- {
        rev = append(rev, s[i])
    }

    for i := k; i < len(s); i++ {
        rev = append(rev, s[i])
    }

    return string(rev)
}