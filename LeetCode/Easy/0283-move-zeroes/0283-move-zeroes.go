func moveZeroes(nums []int)  {
    n := len(nums)
    for i := n-2; i >= 0; i-- {
        for j := i ; j < n -1; j++ {
            if (nums[j] == 0 && nums[j+1] != 0) {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}