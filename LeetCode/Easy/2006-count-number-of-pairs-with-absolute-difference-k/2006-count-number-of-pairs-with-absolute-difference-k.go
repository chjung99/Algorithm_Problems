func countKDifference(nums []int, k int) int {
    cnt := 0

    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if max(nums[i]-nums[j], nums[j]-nums[i]) == k {
                cnt ++
            }
        }
    }
    return cnt
}