func minOperations(nums []int, k int) int {
    sum := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        sum += nums[i]
    }

    return sum % k
}