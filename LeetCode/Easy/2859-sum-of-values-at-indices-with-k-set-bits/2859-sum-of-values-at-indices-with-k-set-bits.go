func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if setCount(i, k) {
            sum += nums[i]
        }
    }
    return sum
}

func setCount(x, k int) bool {
    cnt := 0
    for x != 0 {
        if x % 2 == 1 {
            cnt ++
        }
        x = x / 2
    }
    return cnt == k
}