func reverseWords(s string) string {
    tokens := strings.Fields(s)
    n := len(tokens)
    ans := ""

    for i := n - 1; i >= 0; i-- {
        ans += tokens[i]
        if (i != 0) {
            ans += " "
        }
    }
    return ans
}