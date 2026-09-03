func reverseDegree(s string) int {
    out := 0

    for idx, c := range s {
        out += (26 - int(c - 'a')) * (idx + 1)
    }

    return out
}