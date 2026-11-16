func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    cnt := 0
    n := len(nums1)
    m := len(nums2)

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if nums1[i] % (nums2[j] * k) == 0 {
                cnt ++
            }
        }
    }
    return cnt
}