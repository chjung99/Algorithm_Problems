func minPartitions(n string) int {
    ret := 0

    for _, c := range n {
        ret = max(ret, int(c - '0'))
    }
    return ret
}