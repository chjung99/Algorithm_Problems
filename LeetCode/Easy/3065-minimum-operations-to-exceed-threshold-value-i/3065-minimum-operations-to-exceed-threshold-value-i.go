func minOperations(nums []int, k int) int {
    cnt := 0
    for _, num := range nums {
        if num < k {
            cnt += 1
        }
    }
    return cnt
}