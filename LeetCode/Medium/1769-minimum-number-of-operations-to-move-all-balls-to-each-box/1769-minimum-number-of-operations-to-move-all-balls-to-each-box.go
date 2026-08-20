func minOperations(boxes string) []int {
    ret := make([]int, 0)
    n := len(boxes)

    for i := 0; i < n; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if (boxes[j] == '1') {
                cnt += max(j-i, i-j)
            }
        }
        ret = append(ret, cnt)
    }
    return ret
}