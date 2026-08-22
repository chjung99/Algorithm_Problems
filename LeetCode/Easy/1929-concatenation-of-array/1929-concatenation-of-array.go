func getConcatenation(nums []int) []int {
    n := len(nums)
    ans := make([]int, 0)

    for i := 0; i < n; i++ {
        ans = append(ans, nums[i])
    }

    for i := 0; i < n; i++ {
        ans = append(ans, nums[i])
    }

    return ans
}