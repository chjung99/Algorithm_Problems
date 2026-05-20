func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    prefix := make([]int, n) 
    suffix := make([]int, n)

    for i := 0; i < n; i++ {
        if (i == 0) {
            prefix[i] = 1
        } else {
            prefix[i] = prefix[i-1] * nums[i-1]
        }
    }

    for i := n-1; i >= 0; i-- {
        if (i == n-1) {
            suffix[i] = 1
        } else {
            suffix[i] = suffix[i+1] * nums[i+1]
        }
    }

    for i := 0; i < n; i++ {
        ans[i] = prefix[i] * suffix[i]
    }
    return ans
}