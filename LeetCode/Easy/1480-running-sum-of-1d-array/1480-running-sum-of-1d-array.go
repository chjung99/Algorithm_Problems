func runningSum(nums []int) []int {
    accum := nums[0]
    for i := 1; i < len(nums); i++ {
        accum += nums[i]
        nums[i] = accum
    }
    return nums
}