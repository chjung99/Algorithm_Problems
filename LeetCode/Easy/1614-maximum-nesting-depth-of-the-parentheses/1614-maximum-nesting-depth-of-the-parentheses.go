func maxDepth(s string) int {
    st := make([]rune, 0)
    ans := 0
    d := 0
    for _, ch := range s {
        ans = max(ans, d)
        if ch == '(' {
            st = append(st, ch)
            d += 1
        } else if ch == ')' {
            st = st[:len(st)-1]
            d -= 1
        }
    }
    return ans
}