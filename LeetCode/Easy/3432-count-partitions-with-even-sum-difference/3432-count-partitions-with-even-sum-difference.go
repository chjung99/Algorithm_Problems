func countPartitions(nums []int) int {
    cnt := 0
    left := 0
    right := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        right += nums[i]
    }

    for i := 0; i < n - 1; i++ {
        left += nums[i]
        right -= nums[i]

        if (left - right) % 2 == 0 {
            cnt++
        }
    }
    return cnt
}