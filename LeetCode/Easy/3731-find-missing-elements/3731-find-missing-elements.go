func findMissingElements(nums []int) []int {
    miss := make([]int, 0)
    n := len(nums)
    sort.Ints(nums)


    for i := 0; i < n; i++ {
        cur := nums[i]
        if (i + 1 < n) {
            for (cur + 1 < nums[i+1]) {
                miss = append(miss, cur + 1)
                cur += 1
            }
        }
    }

    return miss
}