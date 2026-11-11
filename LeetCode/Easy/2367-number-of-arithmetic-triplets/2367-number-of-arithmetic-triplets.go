func arithmeticTriplets(nums []int, diff int) int {
    n := len(nums)
    cnt := 0

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            for k := j+1; k < n; k++ {
                if nums[j]-nums[i] == diff && nums[k] - nums[j] == diff {
                    cnt ++
                }
            }
        }
    }
    return cnt
}