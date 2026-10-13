func createTargetArray(nums []int, index []int) []int {
    n := len(nums)
    arr := make([]int, n)

    for i := 0; i < n; i++ {
        arr[i] = -1
    }

    for i := 0; i < n; i++ {
        if arr[index[i]] != -1 {
            for j := n - 1; j > index[i]; j -- {
                arr[j] = arr[j-1]
            }
        }
        arr[index[i]] = nums[i]
    }
    return arr
}