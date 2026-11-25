func subarraySum(nums []int) int {
    n := len(nums)
    s := 0
    for i := 0; i < n; i++ {
        start := max(0, i - nums[i])
        for j := start; j <= i; j++ {
            s += nums[j]
        }
    }
    return s
}