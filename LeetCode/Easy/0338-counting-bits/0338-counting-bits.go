func countBits(n int) []int {
    ans := make([]int, 0)
    for i := 0; i < n + 1; i ++ {
        ans = append(ans, countSetBit(i))
    }
    return ans
}

func countSetBit(x int) int {
    cnt := 0
    for (x != 0) {
        if (x % 2 == 1) {
            cnt += 1
        }
        x = x/2
    }
    return cnt
}