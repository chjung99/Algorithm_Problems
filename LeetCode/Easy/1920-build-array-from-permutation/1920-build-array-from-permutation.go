func buildArray(nums []int) []int {
    ans := make([]int, 0)
    n := len(nums)

    for i := 0; i < n; i++ {
        ans = append(ans, nums[nums[i]])
    }
    return ans
}