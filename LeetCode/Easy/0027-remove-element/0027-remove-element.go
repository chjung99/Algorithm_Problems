func removeElement(nums []int, val int) int {
    n := len(nums)
    cnt := 0
    for i := 0; i < n; i++ {
        if (nums[i] == val) {
            nums[i] = 51
            cnt += 1 
        }
    }
    sort.Ints(nums)
    return n - cnt
}