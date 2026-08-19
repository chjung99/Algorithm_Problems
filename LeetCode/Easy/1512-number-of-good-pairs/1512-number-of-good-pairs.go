func numIdenticalPairs(nums []int) int {
    cnt := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if (nums[i] == nums[j]) {
                cnt += 1
            }
        }
    }
    return cnt
}