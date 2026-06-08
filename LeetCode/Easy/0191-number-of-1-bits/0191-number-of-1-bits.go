func hammingWeight(n int) int {
    cnt := 0
    for (n != 0) {
        cnt += n % 2
        n = n / 2
    }
    return cnt
}