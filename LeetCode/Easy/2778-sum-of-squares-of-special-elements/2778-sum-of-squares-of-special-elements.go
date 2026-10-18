func sumOfSquares(nums []int) int {
    sum := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if n % (i+1) == 0 {
            sum += nums[i] * nums[i]
        }
    }
    return sum
}