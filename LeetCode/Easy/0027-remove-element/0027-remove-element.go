func removeElement(nums []int, val int) int {
    cnt := 0
    for _, v := range nums {
        if (v != val) {
            nums[cnt] = v
            cnt += 1
        }
    }
    return cnt
}