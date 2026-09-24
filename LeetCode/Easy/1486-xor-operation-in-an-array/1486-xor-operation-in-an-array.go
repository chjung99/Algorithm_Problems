func xorOperation(n int, start int) int {
    ret := 0

    for i := range n {
        ret ^= start + 2 * i
    }

    return ret
}