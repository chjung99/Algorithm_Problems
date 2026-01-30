func lengthOfLastWord(s string) int {
    l := 0
    for i := len(s) - 1; i >=0; i-- {
        if (l==0 && s[i]==' ') {
            continue
        } else if (l>0 && s[i]==' ') {
            break
        } else {
            l += 1
        }
    }
    return l
}