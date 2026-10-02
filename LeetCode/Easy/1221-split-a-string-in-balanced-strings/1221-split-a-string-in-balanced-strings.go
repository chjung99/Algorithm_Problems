func balancedStringSplit(s string) int {
    lCnt := 0
    rCnt := 0

    cnt := 0

    for i := 0; i < len(s); i ++ {
        if s[i] == 'L' {
            lCnt += 1
        } else {
            rCnt += 1
        }

        if (lCnt == rCnt) {
            lCnt = 0
            rCnt = 0
            cnt += 1
        }
    }
    return cnt
}