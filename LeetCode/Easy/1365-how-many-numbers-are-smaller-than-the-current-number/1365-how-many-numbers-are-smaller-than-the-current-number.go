func smallerNumbersThanCurrent(nums []int) []int {
    ret := make([]int, 0)
    n := len(nums)

    for i := 0; i < n; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            if (nums[i] > nums[j]) {
                cnt += 1
            }
        }
        ret = append(ret, cnt)
    }
    return ret
}