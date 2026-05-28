func rotate(nums []int, k int)  {
    n := len(nums)
    copied := make([]int, n)
    copy(copied, nums)

    for i := 0; i < n; i++ {
        nums[(i+k)% n] = copied[i]
    }
    
}
 