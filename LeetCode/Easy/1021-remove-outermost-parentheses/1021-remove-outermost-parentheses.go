func removeOuterParentheses(s string) string {
    l := 0
    r := 0
    ret := ""

    for _, c := range s {
        if c == '(' {
            l += 1
            if l > 1 {
                ret += string(c)
            }
        } else if c == ')' {
            r += 1
            if l == r {
                l = 0
                r = 0
            } else {
                ret += string(c)
            }
        }
    }
    return ret
}