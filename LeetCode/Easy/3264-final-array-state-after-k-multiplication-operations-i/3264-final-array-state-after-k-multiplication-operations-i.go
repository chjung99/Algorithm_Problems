func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        minValue := nums[0]
        minIdx := 0
        
        for j := 0; j < len(nums); j++ {
            if minValue > nums[j] {
                minValue = nums[j]
                minIdx = j
            }
        }

        nums[minIdx] *= multiplier
    }
    return nums
}