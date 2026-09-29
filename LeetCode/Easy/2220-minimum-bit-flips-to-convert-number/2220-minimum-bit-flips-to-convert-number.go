func minBitFlips(start int, goal int) int {
    cnt := 0
    x := start ^ goal

    for x != 0 {
        if x % 2 == 1 {
            cnt += 1
        }
        x = x / 2
    }

    return cnt
}
