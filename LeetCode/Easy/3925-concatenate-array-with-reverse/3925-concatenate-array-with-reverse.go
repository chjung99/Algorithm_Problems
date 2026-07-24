func concatWithReverse(nums []int) []int {
    arr := make([]int, 0)
    n := len(nums)
    for i := 0; i < n; i++ {
        arr = append(arr, nums[i])
    }

    for i := n-1; i >= 0; i-- {
        arr = append(arr, nums[i])
    }

    return arr
}