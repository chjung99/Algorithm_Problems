func isPalindrome(s string) bool {
    n := len(s)
    var t []byte

    for i := 0; i < n; i++ {
        if (s[i] >= 'A' && s[i] <= 'Z') {
            t = append(t, 'a' + s[i] - 'A')
        } else if (s[i] >= 'a' && s[i] <= 'z') {
            t = append(t, s[i])
        } else if (s[i] >= '0' && s[i] <= '9') {
            t = append(t, s[i])
        }
    }

    k := len(t)

    for i := 0; i < k; i++ {
        if (t[i] != t[k-i-1]) {
            return false
        }
    }
    return true
}