func shuffle(nums []int, n int) []int {
    shuffled := make([]int, 0)

    for i := 0; i < n; i++ {
        shuffled = append(shuffled, nums[i])
        shuffled = append(shuffled, nums[i + n])
    }
    
    return shuffled
}