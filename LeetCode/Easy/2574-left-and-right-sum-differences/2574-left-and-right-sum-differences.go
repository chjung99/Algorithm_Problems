func leftRightDifference(nums []int) []int {
    n := len(nums)

    diff := make([]int, n)
    leftSum := make([]int, n)
    rightSum := make([]int, n)

    for i := 0; i < n; i++ {
        if (i == 0) {
            leftSum[i] = 0
        } else {
            leftSum[i] = leftSum[i-1] + nums[i-1]
        }
    }

    for i := n - 1; i >= 0; i-- {
        if (i == n - 1) {
            rightSum[i] = 0
        } else {
            rightSum[i] = rightSum[i+1] + nums[i+1]
        }
    }

    for i := 0; i < n; i++ {
        diff[i] = max(leftSum[i]-rightSum[i], -leftSum[i]+rightSum[i])
    }

    return diff
}