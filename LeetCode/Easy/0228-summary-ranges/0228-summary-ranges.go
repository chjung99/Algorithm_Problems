func summaryRanges(nums []int) []string {
    var st []int
    var ans []string

    for _, val := range nums {
        if (len(st) == 0) {
            st = append(st, val)
            continue
        }

        if (st[len(st)-1] + 1 == val) {
            st = append(st, val)
            continue
        }

        if (len(st) == 1) {
            ans = append(ans, strconv.Itoa(st[0]))
            st = st[:len(st)-1]
            st = append(st, val)
            continue
        }

        ans = append(ans, strconv.Itoa(st[0]) + "->" + strconv.Itoa(st[len(st)-1]))
        st = st[:0]
        st = append(st, val)
    }
    if (len(st) == 1) {
        ans = append(ans, strconv.Itoa(st[0]))
        st = st[:len(st)-1]
    } else if (len(st) >= 2){
        ans = append(ans, strconv.Itoa(st[0]) + "->" + strconv.Itoa(st[len(st)-1]))
        st = st[:0]
    }
    return ans
}