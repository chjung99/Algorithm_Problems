func sumCounts(nums []int) int {
    n := len(nums)
    ret := 0

    for s := 1; s <= n; s++ {
        for i := 0; i <= n - s; i++ {
            cnt := 0
            m := make(map[int]bool)
            for j := i; j < i+s; j++ {
                _, exist := m[nums[j]]
                if (!exist) {
                    cnt += 1
                    m[nums[j]] = true
                }
            }
            // fmt.Println(i, i+s - 1, cnt)
            ret += cnt * cnt
        }
    }
    return ret
}